package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fernandamelov/device-management-api/app/models"
	"github.com/fernandamelov/device-management-api/app/repository"
	"github.com/fernandamelov/device-management-api/app/utility"
)

func AddDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device

	json.NewDecoder(r.Body).Decode(&device)

	requestJson, err := json.Marshal(device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("[Add device] Request body:", string(requestJson))

	err = utility.ValidateDevice(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.AddDevice(device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("[Add device] Device added with ID:", id)

	device.ID = int(id)

	response := &models.DeviceDeafultResponse{
		Device: &device,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
