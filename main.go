package main

import (
	"github.com/tutorials/go/crud/db"
	"github.com/tutorials/go/crud/handler"

	"log"

	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handler.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books",h.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books",h.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}",h.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}",h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}",h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}