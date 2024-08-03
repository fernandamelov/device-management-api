package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/fernandamelov/device-management-api/app/models"
	"github.com/fernandamelov/device-management-api/app/repository"
	"github.com/fernandamelov/device-management-api/app/router"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "developer")
	os.Setenv("DB_PASSWORD", "golangdeveloper123")
	os.Setenv("DB_NAME", "device_db")

	repository.InitializeDatabase()
	return router.InitializeRouter()
}

func TestAddDevice(t *testing.T) {
	r := setupRouter()

	device := models.Device{
		Name:  "Test Device",
		Brand: "Test Brand",
	}
	deviceJSON, _ := json.Marshal(device)

	req, _ := http.NewRequest("POST", "/devices", bytes.NewBuffer(deviceJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetDevice(t *testing.T) {
	r := setupRouter()

	device := models.Device{
		Name:  "Test Device",
		Brand: "Test Brand",
	}
	repository.AddDevice(device)

	req, _ := http.NewRequest("GET", "/devices/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListDevices(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/devices", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateDevice(t *testing.T) {
	r := setupRouter()

	device := models.Device{
		Name:  "Test Device",
		Brand: "Test Brand",
	}
	id, _ := repository.AddDevice(device)

	updatedDevice := models.Device{
		Name:  "Updated Device",
		Brand: "Updated Brand",
	}
	updatedDeviceJSON, _ := json.Marshal(updatedDevice)

	req, _ := http.NewRequest("PUT", "/devices/"+strconv.FormatInt(id, 10), bytes.NewBuffer(updatedDeviceJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestPartialUpdateDevice(t *testing.T) {
	r := setupRouter()

	device := models.Device{
		Name:  "Test Device",
		Brand: "Test Brand",
	}
	id, _ := repository.AddDevice(device)

	partialUpdate := map[string]interface{}{
		"name": "Partially Updated Device",
	}
	partialUpdateJSON, _ := json.Marshal(partialUpdate)

	req, _ := http.NewRequest("PATCH", "/devices/"+strconv.FormatInt(id, 10), bytes.NewBuffer(partialUpdateJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteDevice(t *testing.T) {
	r := setupRouter()

	device := models.Device{
		Name:  "Test Device",
		Brand: "Test Brand",
	}
	id, _ := repository.AddDevice(device)

	req, _ := http.NewRequest("DELETE", "/devices/"+strconv.FormatInt(id, 10), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestSearchDevicesByBrand(t *testing.T) {
	r := setupRouter()

	device1 := models.Device{
		Name:  "Device1",
		Brand: "BrandA",
	}
	device2 := models.Device{
		Name:  "Device2",
		Brand: "BrandA",
	}
	repository.AddDevice(device1)
	repository.AddDevice(device2)

	req, _ := http.NewRequest("GET", "/devices/brand/BrandA", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
