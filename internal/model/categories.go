package model

type CategoriesModel struct {
	CategoryID  int     `json:"CategoryID" db:"CategoryID"`
	Name        string  `json:"Name" db:"Name"`
	RoomCount   int     `json:"RoomCount" db:"RoomCount"`
	PricePerDay float64 `json:"PricePerDay" db:"PricePerDay"`
}
