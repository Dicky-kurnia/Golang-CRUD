package handler

import (
	"fmt"
	"github.com/tutorials/go/crud/models"
	"encoding/json"
	
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])

	var book models.Book

	if result := h.DB.First(&book, Id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Delete")
}