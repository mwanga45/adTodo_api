package database

import (
	"fmt"
	"os"
	"todo_api_backend/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Databaseconnection() *gorm.DB {
	// load an envirnment variable 
	localhost := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",localhost,user, password, dbname, port  )
     db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	 if err != nil{
		log.Fatalf("Server failed  to connect  with database %d",err)
	 }
	 db.AutoMigrate(&model.Todoinfo{}, &model.User{})
	
    return  db
}
