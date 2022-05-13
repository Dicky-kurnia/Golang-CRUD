package handler

import (
	"github.com/tutorials/go/crud/mocks"
	"github.com/tutorials/go/crud/models"
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"
	"math/rand"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode("Created")
}