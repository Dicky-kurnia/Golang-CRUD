package handler

import (
	"net/http"
	"io/ioutil"
	"github.com/tutorials/go/crud/mocks"
	"encoding/json"
	"github.com/tutorials/go/crud/models"
	"log"
	"strconv"
	"github.com/gorilla/mux"
	
)

func UpdateBook(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body,err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateBook models.Book
	json.Unmarshal(body, &updateBook)

	for index, book := range mocks.Books {
		if book.Id == Id {
			book.Name = updateBook.Name
			book.IsActive = updateBook.IsActive

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "applicaction/json")
			json.NewEncoder(w).Encode("Updated")
		}
	}
}