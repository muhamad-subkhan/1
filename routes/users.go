package routes

import (
	"BE-foodways/pkg/mysql"
	"BE-foodways/repositories"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	UserRepositories := repositories.RepositoriesUser(mysql.DB)
	h := handlers.HandlerUser(UserRepositories)
	

	r.HandleFunc("/user", h.handlers.CreateUser).Methods("POST")
}