package routes

import (
	"BE-foodways/handlers"
	"BE-foodways/pkg/middleware"
	"BE-foodways/pkg/mysql"
	"BE-foodways/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)
  
	r.HandleFunc("/products", middleware.Auth(h.FindProducts)).Methods("GET") //get all
	r.HandleFunc("/products/{id}", h.GetProduct).Methods("GET") //get by id
	r.HandleFunc("/product/{id}", h.ProductDetail).Methods("GET") //get detail by id
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST") //Create


}