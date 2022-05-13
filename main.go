package main

import (
	"github.com/tutorials/go/crud/handler"

	"log"

	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books",handler.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books",handler.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}",handler.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}",handler.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}",handler.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}