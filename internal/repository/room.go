package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const RoomTable = "Rooms"

type RoomRepo struct {
	db *sql.DB
}

func NewRoom(db *sql.DB) *RoomRepo {
	return &RoomRepo{db: db}
}

func (r *RoomRepo) Create(room model.RoomModel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s ( CategoryID, EmployeeID, Capacity,Description) VALUES ($1,$2,$3,$4)", RoomTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}

	res, err := stmt.Exec(room.CategoryID, room.EmployeeID, room.Capacity, room.Description)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(id), nil
}

func (r *RoomRepo) GetAll() ([]model.RoomModel, error) {
	var rooms []model.RoomModel
	query := fmt.Sprintf("SELECT * FROM %s", RoomTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		room := model.RoomModel{}
		err := rows.Scan(&room.RoomID, &room.CategoryID, &room.EmployeeID, &room.Capacity, &room.Description)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)

	}
	return rooms, nil
}

func (r *RoomRepo) GetOne(roomId string) (model.RoomModel, error) {
	var room model.RoomModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE RoomID=$1", RoomTable)
	row := r.db.QueryRow(query, roomId)
	err := row.Scan(&room.RoomID, &room.CategoryID, &room.EmployeeID, &room.Capacity, &room.Description)
	if err != nil {
		return room, err
	}
	return room, nil
}

func (r *RoomRepo) Delete(roomId string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE RoomID=$1", RoomTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(roomId)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
