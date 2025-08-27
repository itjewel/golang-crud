package routes

import (
	"net/http"

	"golang-crud/controllers"
)

func CategoryRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /categories", controllers.GetCategories)        // GET
	mux.HandleFunc("POST /categories/add", controllers.AddCategory)     // POST
	mux.HandleFunc("PUT /categories/update", controllers.UpateCategory) // PUT
	mux.HandleFunc("DELETE /categories/delete", controllers.DeleteCategory)
}
