package controllers

import (
	"encoding/json"
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

func getCats(w http.ResponseWriter, r *http.Request){
	rows, err := database.DB.Query("SELECT id, name from categories")
	if err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}
defer rows.Close()

	var catObject []models.Category
	for rows.Next(){
		var c models.Category
		if error :=rows.Scan(&c.ID, &c.Name); error != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		catObject = append(catObject, c)
	}
	if err := rows.Err(); err != nil{
		http.Error(w, "rows Iteration Errorserver error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
  if error := json.NewEncoder(w).Encode(catObject); error != nil {
	http.Error(w, "jsonEncode Error", http.StatusInternalServerError)
	return
  }
}

func AddCategory(w http.ResponseWriter, r *http.Request) {

	var c models.Category
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("INSERT INTO categories (name) VALUES (?)", c.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	c.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}
