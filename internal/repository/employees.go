package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const EmployTable = "Employees"

type EmployeesRepo struct {
	db *sql.DB
}

func NewEmployees(db *sql.DB) *EmployeesRepo {
	return &EmployeesRepo{db: db}
}

func (r *EmployeesRepo) Create(employee model.EmployeesModel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (FirstName,LastName,Patronymic,BirthDate,PositionID) VALUES ($1,$2,$3,$4,$5)", EmployTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(employee.Firstname, employee.Lastname, employee.Patronymic, employee.BirthDate, employee.PositionID)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(id), nil
}

func (r *EmployeesRepo) GetAll() ([]model.EmployeesModel, error) {
	var employees []model.EmployeesModel
	query := fmt.Sprintf("SELECT * FROM %s", EmployTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var employee model.EmployeesModel
		err := rows.Scan(&employee.EmployeeID, &employee.Firstname, &employee.Lastname, &employee.Patronymic, &employee.BirthDate, &employee.PositionID)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *EmployeesRepo) GetOne(employeeId string) (model.EmployeesModel, error) {
	var employee model.EmployeesModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE EmployeeID=$1", EmployTable)
	row := r.db.QueryRow(query, employeeId)
	err := row.Scan(&employee.EmployeeID, &employee.Firstname, &employee.Lastname, &employee.Patronymic, &employee.BirthDate, &employee.PositionID)
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (r *EmployeesRepo) Delete(employeeId string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE EmployeeID=$1", EmployTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(employeeId)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
