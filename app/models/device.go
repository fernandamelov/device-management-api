package models

type Device struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Brand        string `json:"brand,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
}

type DeviceDeafultResponse struct {
	Message string   `json:"message,omitempty"`
	ID      int64    `json:"id,omitempty"`
	Erro    string   `json:"erro,omitempty"`
	Device  *Device  `json:"device,omitempty"`
	Devices []Device `json:"devices,omitempty"`
}
