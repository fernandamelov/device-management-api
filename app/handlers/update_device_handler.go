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

func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	params := mux.Vars(r)
	id := params["id"]

	log.Println("[Update device] Requested ID:", id)

	err := utility.ValidateID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&device)

	err = utility.ValidateDevice(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.UpdateDevice(deviceID, device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PartialUpdateDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	params := mux.Vars(r)
	id := params["id"]

	log.Println("[Partial update device] Requested ID:", id)

	err := utility.ValidateID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&device)

	if device.Name == "" && device.Brand == "" {
		http.Error(w, "no fields to update", http.StatusBadRequest)
		return
	}

	err = repository.PartialUpdateDevice(deviceID, device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
