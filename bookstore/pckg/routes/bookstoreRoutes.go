package routes

import (
	"github.com/gorilla/mux"
	"github.com/schraderbachar/bookstore/pckg/controllers"
)

var RegisterBookStoreRotues = func(router *mux.Router) {
	router.HanldeFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HanldeFunc("/book/", controllers.GetBook).Methods("GET")
	router.HanldeFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HanldeFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HanldeFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
