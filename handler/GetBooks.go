package handler

import (
	"github.com/tutorials/go/crud/mocks"
	"encoding/json"
	"net/http"
)


func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}