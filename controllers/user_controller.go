package controllers

import (
	"encoding/json"
	"golang-crud/database"
	"golang-crud/models"
	"golang-crud/service"
	"io/ioutil"
	"log"
	"net/http"
)

type UserControllerService struct {
	Service service.UserService
}

func (uc *UserControllerService) AddUser(w http.ResponseWriter, r *http.Request) {
	var u models.Users
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Requested input not valid", http.StatusInternalServerError)
		return
	}
	response, err := uc.Service.AddUser(u)

	if err != nil {
		http.Error(w, "Not Inserted", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (uc *UserControllerService) BulkUpload(w http.ResponseWriter, r *http.Request) {

	getFile, err := ioutil.ReadFile("utls/user.json")
	if err != nil {
		log.Println("Byte problem Right way")
		return
	}

	var users []models.Users
	if err := json.Unmarshal(getFile, &users); err != nil {
		log.Println("Byte problem Right way")
		return
	}

	for _, value := range users {
		_, err := database.DB.Exec("INSERT INTO users (username,email,password,address) VALUES (?,?,?,?)", value.Name, value.Email, value.Password, value.Address)
		if err != nil {
			log.Println("Insert error for user:", value.Email, err, "jwel")
			continue
		}
	}
	response := map[string]interface{}{
		"data": "Successfully inserted",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (uc *UserControllerService) GeAllUser(w http.ResponseWriter, r *http.Request) {
	res, err := uc.Service.GetUsers()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
