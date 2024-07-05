package model

import "time"

type EmployeesModel struct {
	EmployeeID int       `json:"EmployeeID" db:"EmployeeID"`
	Lastname   string    `json:"Lastname" db:"Lastname"`
	Firstname  string    `json:"Firstname" db:"Firstname"`
	Patronymic string    `json:"Patronymic" db:"Patronymic"`
	BirthDate  time.Time `json:"BirthDate" db:"BirthDate"`
	PositionID int       `json:"PositionID" db:"PositionID"`
}
