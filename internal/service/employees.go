package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type EmployeesService struct {
	repo repository.Employees
}

func NewEmployeesService(repo repository.Employees) *EmployeesService {
	return &EmployeesService{repo: repo}
}

func (s *EmployeesService) Create(employees model.EmployeesModel) (string, error) {
	msg, err := s.repo.Create(employees)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *EmployeesService) GetAll() ([]model.EmployeesModel, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *EmployeesService) GetOne(employeeId string) (model.EmployeesModel, error) {
	data, err := s.repo.GetOne(employeeId)
	if err != nil {
		return model.EmployeesModel{}, err
	}
	return data, nil
}

func (s *EmployeesService) Delete(employeeId string) (string, error) {
	msg, err := s.repo.Delete(employeeId)
	if err != nil {
		return "", err
	}
	return msg, nil
}
