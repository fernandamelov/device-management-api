package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fernandamelov/device-management-api/app/models"
	"github.com/fernandamelov/device-management-api/app/repository"
)

func ListDevices(w http.ResponseWriter, r *http.Request) {

	log.Println("[List devices] Requested")

	devices, err := repository.ListDevices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &models.DeviceDeafultResponse{
		Devices: devices,
	}

	json.NewEncoder(w).Encode(response)
}
