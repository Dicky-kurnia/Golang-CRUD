package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"strconv"
	"github.com/tutorials/go/crud/mocks"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, book := range mocks.Books {
		if book.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
		}
	}
}