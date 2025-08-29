package repository

import (
	"golang-crud/database"
	"golang-crud/models"
)

type UserRepository struct{}

func (u *UserRepository) Insert(p models.Users) (int64, error) {
	res, err := database.DB.Exec("INSERT INTO users (username,email,password,address) VALUES (?,?,?,?)", p.Name, p.Email, p.Password, p.Address)
	if err != nil {
		return 0, nil
	}
	lastID, err := res.LastInsertId()
	return lastID, err
}
