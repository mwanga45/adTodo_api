package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func databaseconnection() *gorm.DB {
	// load an envirnment variable 
	localhost := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",localhost,user, password, dbname  )
     db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	 if err != nil{
		log.Fatal("Server failed  to connect  with database")
	 }
    return  db
}
