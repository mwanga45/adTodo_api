package repository

import (
	database "todo_api_backend/db"
	"todo_api_backend/model"
)



func CreateTodo (todo *model.Todoinfo)error{
	db := database.Databaseconnection()
  return   db.First(&todo).Error
}