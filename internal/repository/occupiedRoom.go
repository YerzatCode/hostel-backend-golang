package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const OccupiedRoomTable = "OccupiedRooms" //

type OccupiedRoomsRepo struct {
	db *sql.DB
}

func NewOccupiedRoom(db *sql.DB) *OccupiedRoomsRepo {
	return &OccupiedRoomsRepo{
		db: db,
	}
}

func (r *OccupiedRoomsRepo) Create(occupiedRoom model.OccupiedRoomsModel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (RoomID, GuestID,ServiceID,CheckInDate,CheckOutDate,AccommodationCost) VALUES ($1, $2, $3, $4, $5, $6)", OccupiedRoomTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	result, err := stmt.Exec(occupiedRoom.RoomID, occupiedRoom.GuestID, occupiedRoom.ServiceID, occupiedRoom.CheckInDate, occupiedRoom.CheckOutDate, occupiedRoom.AccommodationCost)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}

func (r *OccupiedRoomsRepo) GetAll() ([]model.OccupiedRoomsModel, error) {
	var rooms []model.OccupiedRoomsModel
	query := fmt.Sprintf("SELECT * FROM %s", OccupiedRoomTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		r := model.OccupiedRoomsModel{}
		err := rows.Scan(&r.ID, &r.RoomID, &r.GuestID, &r.ServiceID, &r.CheckInDate, &r.CheckOutDate, &r.AccommodationCost)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, r)

	}
	return rooms, nil
}

func (r *OccupiedRoomsRepo) GetOne(ORID string) (model.OccupiedRoomsModel, error) {
	var OR model.OccupiedRoomsModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE ID=$1", OccupiedRoomTable)
	row := r.db.QueryRow(query, ORID)
	err := row.Scan(&OR.ID, &OR.RoomID, &OR.GuestID, &OR.ServiceID, &OR.CheckInDate, &OR.CheckOutDate, &OR.AccommodationCost)
	if err != nil {
		return OR, err
	}
	return OR, nil
}

func (r *OccupiedRoomsRepo) Delete(ORID string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE ID=$1", OccupiedRoomTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(ORID)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
