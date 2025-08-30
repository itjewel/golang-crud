package controllers

import (
	"encoding/json"
	"golang-crud/database"
	"golang-crud/models"
	"golang-crud/service"
	"log"
	"net/http"
	"os"
	"strconv"
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

	getFile, err := os.ReadFile("utls/user.json")
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"data": "Succfullay Inserted"})
}

func (uc *UserControllerService) GeAllUser(w http.ResponseWriter, r *http.Request) {
	res, err := uc.Service.GetUsers()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func (uc *UserControllerService) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	conUserIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
	}

	res, err := uc.Service.GetUser(conUserIdInt)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
func (uc *UserControllerService) GetTextSearch(w http.ResponseWriter, r *http.Request) {
	// userId := r.URL.Query().Get("user_id")
	var c models.Users
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Println("Input error")
		return
	}

	res, err := uc.Service.GetTextSearch(c)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func (uc *UserControllerService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req models.Users
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Input not valid", http.StatusInternalServerError)
		return
	}

	res, err := uc.Service.UpdateUser(ctx, req)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{
		"items": res,
	}, "msg": "successfuly insert"})

}

func (uc *UserControllerService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var c models.Users
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Println("Input error")
		return
	}
	err := uc.Service.DeleteUser(c)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "delete succes"})

}
