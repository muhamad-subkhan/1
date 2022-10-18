package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fmt.Print("Server running on localhost : 5000")
	http.ListenAndServe("localhost: 5000", r)

}