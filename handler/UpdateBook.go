package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/tutorials/go/crud/models"
	"log"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
	
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body,err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateBook models.Book
	json.Unmarshal(body, &updateBook)

	var book models.Book

	if result := h.DB.First(&book, Id);result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Name = updateBook.Name
	book.IsActive = updateBook.IsActive

	h.DB.Save(&book)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "applicaction/json")
	json.NewEncoder(w).Encode("Updated")
}