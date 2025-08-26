package routes

import (
	"net/http"

	"golang-crud/controllers"
)

func CategoryRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/categories", controllers.GetCategories)   // GET
	mux.HandleFunc("/categories/add", controllers.AddCategory) // POST
}
