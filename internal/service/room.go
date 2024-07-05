package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type RoomService struct {
	repo repository.Rooms
}

func NewRoomService(repo repository.Rooms) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) Create(room model.RoomModel) (string, error) {
	msg, err := s.repo.Create(room)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *RoomService) GetAll() ([]model.RoomModel, error) {
	data, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *RoomService) GetOne(roomId string) (model.RoomModel, error) {
	data, err := s.repo.GetOne(roomId)

	if err != nil {
		return model.RoomModel{}, err
	}

	return data, nil
}

func (s *RoomService) Delete(roomId string) (string, error) {
	msg, err := s.repo.Delete(roomId)
	if err != nil {
		return "", err
	}
	return msg, nil
}
