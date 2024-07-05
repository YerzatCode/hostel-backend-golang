package model

import "time"

type OccupiedRoomsModel struct {
	ID                int       `json:"id"`
	RoomID            int       `json:"RoomID"`
	GuestID           int       `json:"GuestID"`
	ServiceID         int       `json:"ServiceID"`
	CheckInDate       time.Time `json:"CheckInDate"`
	CheckOutDate      time.Time `json:"CheckOutDate"`
	AccommodationCost float64   `json:"AccommodationCost"`
}
