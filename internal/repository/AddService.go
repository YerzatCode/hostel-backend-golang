package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const AddServiceTable = "AdditionalServices"

type AdditionalServicesRepo struct {
	db *sql.DB
}

func NewAdditionalServices(db *sql.DB) *AdditionalServicesRepo {
	return &AdditionalServicesRepo{
		db: db,
	}
}

func (r *AdditionalServicesRepo) Create(AddModel model.AdditionalServicesModel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (ServiceName, Price) VALUES ($1, $2)", AddServiceTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec(AddModel.ServiceName, AddModel.Price)
	if err != nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}

func (r *AdditionalServicesRepo) GetAll() ([]model.AdditionalServicesModel, error) {
	var services []model.AdditionalServicesModel
	query := fmt.Sprintf("SELECT * FROM %s", AddServiceTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		service := model.AdditionalServicesModel{}
		err := rows.Scan(&service.ServiceID, &service.ServiceName, &service.Price)
		if err != nil {
			return nil, err
		}
		services = append(services, service)

	}
	return services, nil
}

func (r *AdditionalServicesRepo) GetOne(ServiceID string) (model.AdditionalServicesModel, error) {
	var service model.AdditionalServicesModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE ServiceID=$1", AddServiceTable)
	row := r.db.QueryRow(query, ServiceID)
	err := row.Scan(&service.ServiceID, &service.ServiceName, &service.Price)
	if err != nil {
		return service, err
	}
	return service, nil
}

func (r *AdditionalServicesRepo) Delete(ServiceID string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE ServiceID=$1", AddServiceTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(ServiceID)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
