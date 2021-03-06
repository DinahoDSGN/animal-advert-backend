package services

import (
	"petcard/pkg/models"
	"petcard/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetList(id int) (models.User, error) {
	return s.repo.GetList(id)
}

func (s *UserService) Delete(id int) (models.User, error) {
	return s.repo.Delete(id)
}

func (s *UserService) UpdateRating(id int, data models.User) (float32, error) {
	return s.repo.UpdateRating(id, data)
}

func (s *UserService) Update(id int, data models.User) (models.User, error) {
	return s.repo.Update(id, data)
}
