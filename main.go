package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vairarchi/SimpleRESTApi/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
