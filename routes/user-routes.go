package routes

import (
	"golang-crud/controllers"
	"golang-crud/repository"
	"golang-crud/service"
	"net/http"
)

func UserRoutes(mux *http.ServeMux) {
	repository := repository.UserRepository{}
	service := service.UserService{Repo: repository}
	handler := controllers.UserControllerService{Service: service}
	mux.HandleFunc("POST /user-add", handler.AddUser)
	mux.HandleFunc("GET /bulk-user", handler.BulkUpload)

}
