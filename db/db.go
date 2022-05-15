package db


import (
	"github.com/tutorials/go/crud/models"
    "log"


    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
    dbURL := "postgres://postgres:postgres@localhost:5432"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Book{})

    return db
}