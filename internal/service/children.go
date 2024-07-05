package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type ChildrenService struct {
	repo repository.Children
}

func NewChildrenService(repo repository.Children) *ChildrenService {
	return &ChildrenService{repo: repo}
}

func (s *ChildrenService) Create(chil model.ChildrenModel) (string, error) {
	msg, err := s.repo.Create(chil)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *ChildrenService) GetAll() ([]model.ChildrenModel, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *ChildrenService) GetOne(childrenId string) (model.ChildrenModel, error) {
	data, err := s.repo.GetOne(childrenId)
	if err != nil {
		return model.ChildrenModel{}, err
	}
	return data, nil

}

func (s *ChildrenService) Delete(childrenId string) (string, error) {
	msg, err := s.repo.Delete(childrenId)
	if err != nil {
		return "", nil
	}
	return msg, nil
}
