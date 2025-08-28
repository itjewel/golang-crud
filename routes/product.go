package routes

import (
	"golang-crud/controllers"
	"golang-crud/repository"
	"golang-crud/service"
	"net/http"
)

func ProductRoutes(mux *http.ServeMux) {
	getRepo := &repository.ProductRepository{}
	getService := &service.ProductService{Repo: getRepo}
	getController := &controllers.ControllerService{Service: getService}
	mux.HandleFunc("GET /get-product", getController.GetProductInfo)
	mux.HandleFunc("POST /add-product", getController.AddProduct)

}
