package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const PositionTable = "Positions"

type PositionRepo struct {
	db *sql.DB
}

func NewPosition(db *sql.DB) *PositionRepo {
	return &PositionRepo{db: db}
}

func (r *PositionRepo) Create(pos model.PositionModel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (Title,Salary) VALUES ($1,$2)", PositionTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec(pos.Title, pos.Salary)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(id), nil
}

func (r *PositionRepo) GetAll() ([]model.PositionModel, error) {
	var positions []model.PositionModel
	query := fmt.Sprintf("SELECT * FROM %s", PositionTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var pos model.PositionModel
		err := rows.Scan(&pos.PositionID, &pos.Title, &pos.Salary)
		if err != nil {
			return nil, err
		}
		positions = append(positions, pos)
	}
	return positions, nil
}

func (r *PositionRepo) GetOne(posId string) (model.PositionModel, error) {
	var pos model.PositionModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE PositionID=$1", PositionTable)
	row := r.db.QueryRow(query, posId)
	err := row.Scan(&pos.PositionID, &pos.Title, &pos.Salary)
	if err != nil {
		return pos, err
	}
	return pos, nil
}

func (r *PositionRepo) Delete(posId string) (string, error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE PositionID=$1", PositionTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(posId)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
