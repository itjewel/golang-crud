package routes

import (
	"net/http"

	"golang-crud/controllers"
	"golang-crud/repository"
	"golang-crud/service"
)

func CategoryRoutes(mux *http.ServeMux) {
	categoryRepo := &repository.CategoryRepository{}
	categoryService := &service.CategoryService{Repo: categoryRepo}
	categoryController := &controllers.CategoryController{Service: categoryService}

	mux.HandleFunc("GET /categories", categoryController.GetCategories) // GET
	mux.HandleFunc("GET /category/bulk-upload", categoryController.BulkUpload)
	// mux.HandleFunc("GET /categories-all", controllers.GetAllItem)       // GET
	// mux.HandleFunc("GET /categories-one", controllers.GetOneItem)       // GET
	// mux.HandleFunc("POST /categories/add", controllers.AddCategory)     // POST
	// mux.HandleFunc("PUT /categories/update", controllers.UpateCategory) // PUT
	// mux.HandleFunc("DELETE /categories/delete", controllers.DeleteCategory)
	// mux.HandleFunc("GET /categories/like", controllers.GetLike)
	// mux.HandleFunc("GET /categories/range", controllers.GetRange)
	// mux.HandleFunc("GET /categories/sort", controllers.GetSort)
}
