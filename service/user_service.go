package service

import (
	"errors"
	"golang-crud/models"
	"golang-crud/repository"
	"log"
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

func (s *UserService) GetUsers() ([]models.Users, error) {
	res, err := s.Repo.GetAll()
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	return res, nil
}

func (s *UserService) GetUser(userId int) (*models.Users, error) {
	res, err := s.Repo.GetOneUser(userId)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	return res, nil
}
