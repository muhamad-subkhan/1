package main

import (
	"BE-foodways/database"
	"BE-foodways/pkg/mysql"
	"BE-foodways/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mysql.Database()

	database.RunMigration()

	r := mux.NewRouter()

	routes.Routes(r.PathPrefix("/api/v1").Subrouter())

	fmt.Print("Server running on localhost : 5000")
	http.ListenAndServe("localhost: 5000", r)

}
