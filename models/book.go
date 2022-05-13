package models

import (
	"time"
)

type Book struct {
	Id    int `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	IsActive bool 
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}


