package main

import (
	"BE-foodways/database"
	"BE-foodways/pkg/mysql"
	"BE-foodways/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv" // import this package
)

func main() {

	
	errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    }

	mysql.Database()

	database.RunMigration()

	r := mux.NewRouter()

	routes.Routes(r.PathPrefix("/api/v1").Subrouter())

	fmt.Print("Server running on localhost : 5000")
	http.ListenAndServe("localhost: 5000", r)

}
