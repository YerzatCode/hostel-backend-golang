package model

import "time"

type ChildrenModel struct {
	ChildID    int       `json:"ChildID"`
	GuestID    int       `json:"GuestID"`
	Lastname   string    `json:"Lastname"`
	Firstname  string    `json:"Firstname"`
	Patronymic string    `json:"Patronymic"`
	BirthDate  time.Time `json:"BirthDate"`
}
