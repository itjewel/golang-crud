package routes

import (
	"net/http"

	"golang-crud/controllers"
)

func CategoryRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /categories", controllers.GetCategories)        // GET
	mux.HandleFunc("GET /categories-all", controllers.GetAllItem)       // GET
	mux.HandleFunc("GET /categories-one", controllers.GetOneItem)       // GET
	mux.HandleFunc("POST /categories/add", controllers.AddCategory)     // POST
	mux.HandleFunc("PUT /categories/update", controllers.UpateCategory) // PUT
	mux.HandleFunc("DELETE /categories/delete", controllers.DeleteCategory)
	mux.HandleFunc("GET /categories/like", controllers.GetLike)
	mux.HandleFunc("GET /categories/range", controllers.GetRange)
	mux.HandleFunc("GET /categories/sort", controllers.GetSort)
}
