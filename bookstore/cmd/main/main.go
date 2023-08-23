package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/schraderbachar/bookstore/pckg/routes"
)

// create server, and tell where the routes are
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal("ListenAndServe: ", http.ListenAndServe("localhost:9010", r)) // listen on port 9010
}
