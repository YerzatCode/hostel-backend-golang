package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type OccupiedService struct {
	repo repository.OccupiedRooms
}

func NewOccupiedService(repo repository.OccupiedRooms) *OccupiedService {
	return &OccupiedService{repo: repo}
}

func (s *OccupiedService) Create(occupiedRoom model.OccupiedRoomsModel) (string, error) {
	msg, err := s.repo.Create(occupiedRoom)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *OccupiedService) GetAll() ([]model.OccupiedRoomsModel, error) {
	data, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *OccupiedService) GetOne(occupiedRoomId string) (model.OccupiedRoomsModel, error) {
	data, err := s.repo.GetOne(occupiedRoomId)

	if err != nil {
		return model.OccupiedRoomsModel{}, err
	}

	return data, nil
}

func (s *OccupiedService) Delete(occupiedRoomId string) (string, error) {
	msg, err := s.repo.Delete(occupiedRoomId)
	if err != nil {
		return "", err
	}
	return msg, err
}
