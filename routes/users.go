package routes

import (
	"BE-foodways/pkg/mysql"
	"BE-foodways/handlers"
	"BE-foodways/repositories"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	UserRepository := repositories.RepositoriesUser(mysql.DB)

	h := handlers.HandlerUser(UserRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
}