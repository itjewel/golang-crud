package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-crud/database"
	"golang-crud/models"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		rows.Scan(&c.ID, &c.Name)
		categories = append(categories, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func UpateCategory(w http.ResponseWriter, r *http.Request) {
	var req models.Category
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "response Error", http.StatusInternalServerError)
		return
	}

	result, err := database.DB.Exec("UPDATE categories SET name = ? WHERE id = ? ", req.Name, req.ID)
	if err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check rows affected", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No category found with given ID", http.StatusNotFound)
		return
	}
	resMessage := models.Response{
		Message: "Success Update",
		Status:  200,
		Data:    req,
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resMessage)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {

	var c models.CategoryCreateRequest
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := database.DB.Exec("INSERT INTO categories (name) VALUES (?)", c.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := response.LastInsertId()

	fmt.Println(response, id)
	customeRes := models.Category{
		ID:   int(id),
		Name: c.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customeRes)
}
