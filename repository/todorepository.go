package repository

import (
	database "todo_api_backend/db"
	"todo_api_backend/model"
)

func CreateTodo(todo *model.Todoinfo) error {
	db := database.Databaseconnection()
	return db.Create(&todo).Error
}

func Gettodolist()([]model.Todoinfo,error){
	db := database.Databaseconnection()
	var todo []model.Todoinfo
	err := db.Find(&todo).Error
	if err != nil{
		return  nil , err
	}
	return todo,nil
}

