package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const ChildrenTable = "Children"

type ChildrenRepo struct {
	db *sql.DB
}

func NewChildren(db *sql.DB) *ChildrenRepo {
	return &ChildrenRepo{db: db}
}

func (r *ChildrenRepo) Create(children model.ChildrenModel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (GuestID,Lastname,Firstname,Patronymic,BirthDate) VALUES ($1,$2,$3,$4,$5)", ChildrenTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(children.GuestID, children.Lastname, children.Lastname, children.Firstname, children.Patronymic, children.BirthDate)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}

func (r *ChildrenRepo) GetAll() ([]model.ChildrenModel, error) {
	var childrens []model.ChildrenModel
	query := fmt.Sprintf("SELECT * FROM %s", ChildrenTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		children := model.ChildrenModel{}
		err := rows.Scan(&children.ChildID, &children.GuestID, &children.Lastname, &children.Firstname, &children.Patronymic, &children.BirthDate)
		if err != nil {
			return nil, err
		}
		childrens = append(childrens, children)

	}
	return childrens, nil
}

func (r *ChildrenRepo) GetOne(id string) (model.ChildrenModel, error) {
	var children model.ChildrenModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE ChildID = $1", ChildrenTable)
	row := r.db.QueryRow(query, id)
	err := row.Scan(&children.ChildID, &children.GuestID, &children.Lastname, &children.Firstname, &children.Patronymic, &children.BirthDate)
	if err != nil {
		return children, err
	}
	return children, nil
}

func (r *ChildrenRepo) Delete(childrenId string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE ChildID=$1", ChildrenTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(childrenId)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
