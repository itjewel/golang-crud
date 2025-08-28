package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang-crud/database"
	"golang-crud/models"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name,price FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Price); err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
			continue
		}
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

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("catId") // string
	if idStr == "" {
		http.Error(w, "Missing catId", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr) // string -> int conversion
	if err != nil {
		http.Error(w, "Invalid catId", http.StatusBadRequest)
		return
	}
	result, err := database.DB.Exec("DELETE FROM categories WHERE id = ? ", id)
	if err != nil {
		log.Println("sql error", err)
		http.Error(w, "sql Error", http.StatusInternalServerError)
		return
	}

	rowErrectFields, _ := result.RowsAffected()
	if rowErrectFields == 0 {
		http.Error(w, "No fild updatede", http.StatusNotFound)
		return
	}
	successMessage := models.Response{
		Message: "Success Delete",
		Status:  200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successMessage)

	// log.Println("Error", id)
}

func Jewel(w http.ResponseWriter, r *http.Request) {
	println("hello")
}
func AddCategory(w http.ResponseWriter, r *http.Request) {

	var c models.CategoryCreateRequest
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := database.DB.Exec("INSERT INTO categories (name, price) VALUES (?, ?)", c.Name, c.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := response.LastInsertId()

	customeRes := models.Category{
		ID:    int(id),
		Name:  c.Name,
		Price: float64(c.Price),
	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(customeRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)

}

func GetAllItem(w http.ResponseWriter, r *http.Request) {
	res, err := database.DB.Query("SELECT id,name FROM categories")
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
	}
	defer res.Close()
	var responseData []models.Category
	for res.Next() {
		var c models.Category
		res.Scan(&c.ID, &c.Name)
		responseData = append(responseData, c)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func GetOneItem(w http.ResponseWriter, r *http.Request) {
	var reqData models.Category
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Decode Error", http.StatusInternalServerError)
		return
	}
	var newObject models.Category
	erre := database.DB.QueryRow(
		"SELECT id, name FROM categories WHERE id = ?",
		reqData.ID,
	).Scan(&newObject.ID, &newObject.Name)

	if erre != nil {
		log.Printf("query", erre)
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newObject)
}

func GetLike(w http.ResponseWriter, r *http.Request) {
	var queryParam = r.URL.Query().Get("name")

	res, err := database.DB.Query("SELECT id,name FROM categories WHERE name LIKE ?", "%"+queryParam+"%")

	if err != nil {
		log.Println("db error", err)
		return
	}
	defer res.Close()

	var responseObject []models.Category
	for res.Next() {
		var c models.Category
		if err := res.Scan(&c.ID, &c.Name); err != nil {
			log.Println("db error", err)
			continue
		}
		responseObject = append(responseObject, c)

	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(responseObject)

	if err != nil {
		log.Println("db error", err)
		return
	}
	w.Write(jsonData)
}

func GetRange(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	res, err := database.DB.Query("SELECT id,name,price FROM categories WHERE price BETWEEN ? AND ?", from, to)

	if err != nil {
		log.Println("db error", err)
		return
	}
	defer res.Close()

	var responseObject []models.Category
	for res.Next() {
		var c models.Category
		if err := res.Scan(&c.ID, &c.Name, &c.Price); err != nil {
			log.Println("db error", err)
			continue
		}
		responseObject = append(responseObject, c)

	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(responseObject)

	if err != nil {
		log.Println("db error", err)
		return
	}
	w.Write(jsonData)
}
func GetSort(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Like search")
}
