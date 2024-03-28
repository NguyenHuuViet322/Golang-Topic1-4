package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	var err error

	log.Println("Getting DB")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "123", "testdb")
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	db.AutoMigrate(&todoItem{})

	if err != nil {
		log.Println("Err while getting DB")
		panic(err)
	}
	return db
}
