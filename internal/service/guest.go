package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type GuestService struct {
	repo repository.Guests
}

func NewGuestService(repo repository.Guests) *GuestService {
	return &GuestService{repo: repo}
}

func (s *GuestService) Create(guest model.GuestModel) (string, error) {

	msg, err := s.repo.Create(guest)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *GuestService) GetAll() ([]model.GuestModel, error) {

	data, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *GuestService) GetOne(guestId string) (model.GuestModel, error) {
	data, err := s.repo.GetOne(guestId)

	if err != nil {
		return model.GuestModel{}, err
	}

	return data, nil
}

func (s *GuestService) Delete(guestId string) (string, error) {
	msg, err := s.repo.Delete(guestId)
	if err != nil {
		return "", err
	}
	return msg, err
}

func (s *GuestService) Update(guestID string, guest model.GuestModel) error {
	err := s.repo.Update(guestID, guest)
	if err != nil {
		return err
	}
	return nil
}
