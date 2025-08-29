package controllers

import (
	"encoding/json"
	"golang-crud/database"
	"golang-crud/models"
	"golang-crud/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

/*
func (uc *UserControllerService) BulkUpload(w http.ResponseWriter, r *http.Request) {
	// Open JSON file
	file, err := os.Open("utls/user.json") // adjust path as per run location
	if err != nil {
		http.Error(w, "Cannot open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Cannot read file", http.StatusInternalServerError)
		return
	}

	// Unmarshal JSON into a slice of users
	var users []models.Users
	if err := json.Unmarshal(data, &users); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		return
	}

	// Optionally: save users to DB using service
	// for _, u := range users {
	//     uc.Service.AddUser(u)
	// }

	// Response
	response := map[string]interface{}{
		"status": "success",
		"data":   users, // send proper JSON array
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

*/

func (uc *UserControllerService) BulkUpload(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("utls/user.json")
	if err != nil {
		log.Println("no get Right way")
		return
	}
	defer file.Close()
	getFile, err := ioutil.ReadAll(file)
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
