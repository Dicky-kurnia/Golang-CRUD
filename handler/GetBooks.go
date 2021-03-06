package handler

import (
	"fmt"
	"github.com/tutorials/go/crud/models"
	"github.com/tutorials/go/crud/mocks"
	"encoding/json"
	"net/http"
)


func(h handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	var book []models.Book

	if result := h.DB.Find(&book); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}