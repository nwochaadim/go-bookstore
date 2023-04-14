package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nwochaadim/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	fmt.Println("Listening on port 8080...")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Errorf("Error occurred while starting http server")
	}
}
