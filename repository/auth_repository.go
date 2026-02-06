package repository

import (
	database "todo_api_backend/db"
	"todo_api_backend/model"
)

func CreateUser(user *model.User)error{
  db := database.Databaseconnection()
  return db.Create(&user).Error
}