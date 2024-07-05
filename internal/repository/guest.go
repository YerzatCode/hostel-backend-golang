package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const GuestTable = "Guests"

type GuestRepo struct {
	db *sql.DB
}

func NewGuest(db *sql.DB) *GuestRepo {
	return &GuestRepo{
		db: db,
	}
}

func (r *GuestRepo) Create(guest model.GuestModel) (string, error) {

	query := fmt.Sprintf("INSERT INTO %s(Lastname, Firstname, Patronymic, PassportNumber, BirthDate, CityOfOrigin) VALUES ($1, $2, $3, $4, $5, $6)", GuestTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}

	res, err := stmt.Exec(guest.Lastname, guest.Firstname, guest.Patronymic, guest.PassportNumber, guest.BirthDate.Format("2006-01-02"), guest.CityOfOrigin)
	if err != nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
func (r *GuestRepo) GetAll() ([]model.GuestModel, error) {
	var quests []model.GuestModel

	rows, err := r.db.Query("SELECT * FROM Guests")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		g := model.GuestModel{}
		err := rows.Scan(&g.GuestID, &g.Lastname, &g.Firstname, &g.Patronymic, &g.PassportNumber, &g.BirthDate, &g.CityOfOrigin)
		if err != nil {
			return nil, err
		}
		quests = append(quests, g)

	}
	return quests, nil
}
func (r *GuestRepo) GetOne(guestId string) (model.GuestModel, error) {
	var g model.GuestModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE GuestID=$1", GuestTable)
	row := r.db.QueryRow(query, guestId)
	err := row.Scan(&g.GuestID, &g.Lastname, &g.Firstname, &g.Patronymic, &g.PassportNumber, &g.BirthDate, &g.CityOfOrigin)
	if err != nil {
		return g, err
	}
	return g, nil
}

func (r *GuestRepo) Delete(guestId string) (string, error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE GuestID=$1", GuestTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(guestId)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}

func (r *GuestRepo) Update(guestId string, guest model.GuestModel) error {
	query := fmt.Sprintf("UPDATE %s SET $1 WHERE GuestID=$2", GuestTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(guestId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	_ = id

	return nil
}
