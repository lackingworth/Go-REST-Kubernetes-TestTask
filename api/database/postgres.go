package database

import (
	"fmt"

	m "github.com/lackingworth/Go-REST-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Dsn string = "host=postgres user=postgres password=pass dbname=books port=5432 sslmode=disable" // For Docker deployment
// var Dsn string = "host=0.0.0.0 user=postgres password=pass dbname=books port=5432 sslmode=disable" // For local testing

func NewConnection(dsn string) (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return db, err
	}

	err = db.AutoMigrate(&m.Books{})

	if err != nil {
		fmt.Println(err)
		return db, err
	}
	
	return db, nil
}