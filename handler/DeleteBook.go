package handler

import (
	"encoding/json"
	"github.com/tutorials/go/crud/mocks"
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])

	for index, book := range mocks.Books {
		if book.Id == Id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Delete")
			break
		}
	}
}