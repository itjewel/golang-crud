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
	mux.HandleFunc("GET /users/all-users", handler.GeAllUser)
	mux.HandleFunc("GET /users/one", handler.GetUser)
	mux.HandleFunc("POST /users/textsearch", handler.GetTextSearch)
	mux.HandleFunc("POST /users/pagination", handler.GetUser)
	mux.HandleFunc("POST /users/delete", handler.DeleteUser)
	mux.HandleFunc("PUT /users/update", handler.UpdateUser)

}
