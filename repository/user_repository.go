package repository

import (
	"context"
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

func (u *UserRepository) GetOneUser(userId int) (*models.Users, error) {
	var c models.Users

	err := database.DB.QueryRow("SELECT id,username,email,address,password FROM users WHERE id = ?", userId).Scan(&c.Id, &c.Name, &c.Email, &c.Address, &c.Password)
	if err != nil {
		log.Println(err, "jewel")
		return nil, err
	}

	return &c, nil
}

func (u *UserRepository) TextSearch(req models.Users) ([]models.Users, error) {
	// var c models.Users
	res, err := database.DB.Query("SELECT id,username,email,address,password FROM users WHERE username LIKE ? OR email LIKE ? OR address LIKE ?  ORDER BY id DESC LIMIT 8 OFFSET 0", "%"+req.Name+"%", req.Name+"%", req.Name+"%")
	if err != nil {
		log.Println(err, "jewel")
		return nil, err
	}
	defer res.Close()
	var customeObject []models.Users
	for res.Next() {
		var i models.Users
		if err := res.Scan(&i.Id, &i.Name, &i.Email, &i.Address, &i.Password); err != nil {
			log.Println("error")
			continue
		}
		customeObject = append(customeObject, i)
	}

	return customeObject, nil
}

func (u *UserRepository) Update(ctx context.Context, req models.Users) (int64, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return 0, err
	}
	res, err := tx.ExecContext(ctx,
		"UPDATE users SET username=?, email=?, password=?, address=? WHERE id=?",
		req.Name, req.Email, req.Password, req.Address, req.Id,
	)
	if err != nil {
		log.Println("Update error:", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("RowsAffected error:", err)
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("RowsAffected error:", err)
		return 0, err
	}

	return rowsAffected, nil
}

func (u *UserRepository) DeleteUser(req models.Users) (int64, error) {
	// var c models.Users
	res, err := database.DB.Exec("DELECT FROM users where id =?", req.Id)
	if err != nil {
		log.Println(err, "jewel")
		return 0, err
	}
	id, err := res.RowsAffected()
	if err != nil {
		log.Println(err, "jewel")
		return 0, err
	}

	return id, nil
}
