package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fernandamelov/device-management-api/app/repository"
	"github.com/fernandamelov/device-management-api/app/utility"
	"github.com/gorilla/mux"
)

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	log.Println("[Delete device] Requested ID:", id)

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

	err = repository.DeleteDevice(deviceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
