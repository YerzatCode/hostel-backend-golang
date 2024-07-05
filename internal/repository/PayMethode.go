package repository

import (
	"database/sql"
	"fmt"

	"github.com/YerzatCode/sql_/internal/model"
)

const PayMetTable = "PaymentMethods"

type PayMetRepo struct {
	db *sql.DB
}

func NewPayMet(db *sql.DB) *PayMetRepo {
	return &PayMetRepo{db: db}
}

func (r *PayMetRepo) Create(PayMetModel model.PaymentMethodsModel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (GuestID, PaymentForm,PaidAmount) VALUES ($1,$2,$3)", PayMetTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(PayMetModel.GuestID, PayMetModel.PaymentForm, PayMetModel.PaidAmount)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(id), nil
}

func (r *PayMetRepo) GetAll() ([]model.PaymentMethodsModel, error) {
	var payMets []model.PaymentMethodsModel
	query := fmt.Sprintf("SELECT * FROM %s", PayMetTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		payMet := model.PaymentMethodsModel{}
		err := rows.Scan(&payMet.ID, &payMet.GuestID, &payMet.PaymentForm, &payMet.PaidAmount)
		if err != nil {
			return nil, err
		}
		payMets = append(payMets, payMet)

	}
	return payMets, nil
}

func (r *PayMetRepo) GetOne(id string) (model.PaymentMethodsModel, error) {
	var payMet model.PaymentMethodsModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE ID=$1", PayMetTable)
	row := r.db.QueryRow(query, id)
	err := row.Scan(&payMet.ID, &payMet.GuestID, &payMet.PaymentForm, &payMet.PaidAmount)
	if err != nil {
		return payMet, err
	}
	return payMet, nil
}

func (r *PayMetRepo) Delete(PayId string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE ID=$1", PayMetTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	result, err := stmt.Exec(PayId)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
