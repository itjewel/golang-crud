package service

import (
	"errors"
	"golang-crud/models"
	"golang-crud/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) AddUser(u models.Users) (*models.Users, error) {

	if u.Name == "" {
		return nil, errors.New("name is empty")

	}
	id, err := s.Repo.Insert(u)
	if err != nil {
		return nil, err
	}

	u.Id = int(id)
	return &u, nil

}
