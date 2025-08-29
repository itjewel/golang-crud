package repository

import (
	"golang-crud/database"
	"golang-crud/models"
	"log"
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

func (u *UserRepository) GetAll() ([]models.Users, error) {
	query, err := database.DB.Query("SELECT id,username,email,address,password FROM users")
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	var customObject []models.Users
	for query.Next() {
		var c models.Users
		if err := query.Scan(&c.Id, &c.Name, &c.Email, &c.Address, &c.Password); err != nil {
			log.Println(err)
			return nil, nil
		}
		customObject = append(customObject, c)

	}
	log.Println(customObject)
	return customObject, nil
}
