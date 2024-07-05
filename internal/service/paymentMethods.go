package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type PayMetService struct {
	repo repository.PaymentMethods
}

func NewPayMetService(repo repository.PaymentMethods) *PayMetService {
	return &PayMetService{repo: repo}
}

func (s *PayMetService) Create(paymentMethod model.PaymentMethodsModel) (string, error) {
	msg, err := s.repo.Create(paymentMethod)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *PayMetService) GetAll() ([]model.PaymentMethodsModel, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *PayMetService) GetOne(paymentMethodId string) (model.PaymentMethodsModel, error) {
	data, err := s.repo.GetOne(paymentMethodId)
	if err != nil {
		return model.PaymentMethodsModel{}, err
	}
	return data, nil
}

func (s *PayMetService) Delete(paymentMethodId string) (string, error) {
	msg, err := s.repo.Delete(paymentMethodId)
	if err != nil {
		return "", nil
	}
	return msg, nil
}
