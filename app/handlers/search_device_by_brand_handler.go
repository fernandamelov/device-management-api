package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fernandamelov/device-management-api/app/models"
	"github.com/fernandamelov/device-management-api/app/repository"
	"github.com/fernandamelov/device-management-api/app/utility"
	"github.com/gorilla/mux"
)

func SearchDevicesByBrand(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	brand := params["brand"]

	err := utility.ValidateBrand(brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("[Search devices by brand] Requested brand:", brand)

	devices, err := repository.SearchDevicesByBrand(brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &models.DeviceDeafultResponse{
		Devices: devices,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
