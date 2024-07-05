package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

type CategoriesService struct {
	repo repository.Categories
}

func NewCategoriesService(repo repository.Categories) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (s *CategoriesService) Create(category model.CategoriesModel) (string, error) {
	msg, err := s.repo.Create(category)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *CategoriesService) GetAll() ([]model.CategoriesModel, error) {
	data, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *CategoriesService) GetOne(categoryId string) (model.CategoriesModel, error) {
	data, err := s.repo.GetOne(categoryId)

	if err != nil {
		return model.CategoriesModel{}, err
	}

	return data, nil
}

func (s *CategoriesService) Delete(categoryId string) (string, error) {
	msg, err := s.repo.Delete(categoryId)
	if err != nil {
		return "", err
	}
	return msg, nil
}
