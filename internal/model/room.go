package model

type RoomModel struct {
	RoomID      int    `json:"RoomID" db:"RoomID"`
	CategoryID  int    `json:"CategoryID" db:"CategoryID"`
	EmployeeID  int    `json:"EmployeeID" db:"EmployeeID"`
	Capacity    int    `json:"Capacity" db:"Capacity"`
	Description string `json:"Description" db:"Description"`
}
