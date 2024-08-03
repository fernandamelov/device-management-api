package router

import (
	"github.com/fernandamelov/device-management-api/app/handlers"
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/devices", handlers.AddDevice).Methods("POST")
	r.HandleFunc("/devices/{id}", handlers.GetDeviceByID).Methods("GET")
	r.HandleFunc("/devices", handlers.ListDevices).Methods("GET")
	r.HandleFunc("/devices/{id}", handlers.UpdateDevice).Methods("PUT")
	r.HandleFunc("/devices/{id}", handlers.PartialUpdateDevice).Methods("PATCH")
	r.HandleFunc("/devices/{id}", handlers.DeleteDevice).Methods("DELETE")
	r.HandleFunc("/devices/brand/{brand}", handlers.SearchDevicesByBrand).Methods("GET")
	return r
}
