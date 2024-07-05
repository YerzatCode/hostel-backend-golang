package model

type PositionModel struct {
	PositionID int     `json:"PositionID" db:"PositionID"`
	Title      string  `json:"Title" db:"Title"`
	Salary     float64 `json:"Salary" db:"Salary"`
}
