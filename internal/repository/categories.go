package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const CategoryTable = "Categories"

type CategoryRepo struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) Create(category model.CategoriesModel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (Name,RoomCount,PricePerDay) VALUES($1,$2,$3)", CategoryTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}

	res, err := stmt.Exec(category.Name, category.RoomCount, category.PricePerDay)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), err
}

func (r *CategoryRepo) GetAll() ([]model.CategoriesModel, error) {
	var categories []model.CategoriesModel
	query := fmt.Sprintf("SELECT * FROM %s", CategoryTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := model.CategoriesModel{}
		err := rows.Scan(&c.CategoryID, &c.Name, &c.RoomCount, &c.PricePerDay)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)

	}
	return categories, nil
}

func (r *CategoryRepo) GetOne(categoryId string) (model.CategoriesModel, error) {
	var c model.CategoriesModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE CategoryID=$1", CategoryTable)
	row := r.db.QueryRow(query, categoryId)
	err := row.Scan(&c.CategoryID, &c.Name, &c.RoomCount, &c.PricePerDay)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (r *CategoryRepo) Delete(categoryId string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE CategoryID=$1", CategoryTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(categoryId)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
