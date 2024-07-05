package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type PositionsService struct {
	repo repository.Positions
}

func NewPositionsService(repo repository.Positions) *PositionsService {
	return &PositionsService{repo: repo}
}

func (s *PositionsService) Create(pos model.PositionModel) (string, error) {
	msg, err := s.repo.Create(pos)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *PositionsService) GetAll() ([]model.PositionModel, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *PositionsService) GetOne(posId string) (model.PositionModel, error) {
	data, err := s.repo.GetOne(posId)
	if err != nil {
		return model.PositionModel{}, err
	}
	return data, nil
}

func (s *PositionsService) Delete(posId string) (string, error) {
	msg, err := s.repo.Delete(posId)
	if err != nil {
		return "", err
	}
	return msg, nil
}
