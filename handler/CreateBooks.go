package handler

import (
	"fmt"

	"github.com/tutorials/go/crud/models"
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"
	
)

func (h handler) CreateBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode("Created")
}