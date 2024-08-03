package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fernandamelov/device-management-api/app/models"
	"github.com/fernandamelov/device-management-api/app/repository"
	"github.com/fernandamelov/device-management-api/app/utility"
	"github.com/gorilla/mux"
)

func GetDeviceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("[Get device by ID] Requested ID:", id)

	err := utility.ValidateID(id)
	if err != nil {
		response := &models.DeviceDeafultResponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	deviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := &models.DeviceDeafultResponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	deviceData, err := repository.GetDevice(deviceID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			response := &models.DeviceDeafultResponse{
				Message: "Device not found",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := &models.DeviceDeafultResponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &models.DeviceDeafultResponse{
		Device: &deviceData,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
