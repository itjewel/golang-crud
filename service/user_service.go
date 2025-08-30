package service

import (
	"context"
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

func (s *UserService) GetTextSearch(req models.Users) ([]models.Users, error) {
	res, err := s.Repo.TextSearch(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	return res, nil
}
func (s *UserService) DeleteUser(req models.Users) error {
	if req.Id == 0 {
		log.Println("User id  is empty")
		return nil
	}
	id, err := s.Repo.DeleteUser(req)

	if err != nil {
		log.Println(err)
	}
	if id == 0 {
		log.Println("not deleted")
	}
	// customeObject := map[string]interface{}{
	// 	Id: id,
	// }
	// log.Println(res)
	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, req models.Users) (*models.Users, error) {
	// var object models.Users

	if req.Name == "" {
		log.Println("Name is empty")
		return nil, nil
	}
	id, _ := s.Repo.Update(ctx, req)

	if id == 0 {
		log.Println("Not Update")
		return nil, nil
	}
	// object.Id = int(id)
	// object.Id = 1

	return &req, nil
}
