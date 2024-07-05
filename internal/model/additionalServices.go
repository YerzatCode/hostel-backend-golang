package model

type AdditionalServicesModel struct {
	ServiceID   int     `json:"ServiceID"`
	ServiceName string  `json:"ServiceName"`
	Price       float64 `json:"Price"`
}
