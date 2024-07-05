package model

import "time"

type GuestModel struct {
	GuestID        int       `json:"GuestID" db:"GuestID"`
	Lastname       string    `json:"lastname"`
	Firstname      string    `json:"firstname"`
	Patronymic     string    `json:"patronymic"`
	PassportNumber string    `json:"passportname"`
	BirthDate      time.Time `json:"birthdate"`
	CityOfOrigin   string    `json:"cityoforigin"`
}
