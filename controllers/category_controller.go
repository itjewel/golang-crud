package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-crud/models"
	"golang-crud/service"
)

type CategoryController struct {
	Service *service.CategoryService
}

// GET /categories
func (c *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := c.Service.GetAllCategories(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// POST /categories/add
func (c *CategoryController) AddCategory(w http.ResponseWriter, r *http.Request) {
	var req models.Category
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newCat, err := c.Service.AddCategory(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(newCat)
}

// PUT /categories/update
func (c *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var req models.Category
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.Service.UpdateCategory(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Updated successfully"})
}

// DELETE /categories/delete?catId=1
func (c *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("catId")
	id, _ := strconv.Atoi(idStr)
	if err := c.Service.DeleteCategory(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}
