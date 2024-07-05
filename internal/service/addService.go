package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type AddService struct {
	repo repository.AdditionalServices
}

func NewAddService(repo repository.AdditionalServices) *AddService {
	return &AddService{repo: repo}
}

func (s *AddService) Create(additionalService model.AdditionalServicesModel) (string, error) {
	msg, err := s.repo.Create(additionalService)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *AddService) GetAll() ([]model.AdditionalServicesModel, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *AddService) GetOne(additionalServiceId string) (model.AdditionalServicesModel, error) {
	data, err := s.repo.GetOne(additionalServiceId)
	if err != nil {
		return model.AdditionalServicesModel{}, err
	}
	return data, nil
}

func (s *AddService) Delete(additionalServiceId string) (string, error) {
	msg, err := s.repo.Delete(additionalServiceId)
	if err != nil {
		return "", err
	}
	return msg, nil
}
