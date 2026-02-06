package service

import (
	"errors"
	"todo_api_backend/model"
	"todo_api_backend/repository"
)


func CreateTodo(todo *model.Todoinfo)error{
	if todo.Title == ""{
       return  errors.New("title is requied ")
	}
	return  repository.CreateTodo(todo)
}
func Gettodolist()([]model.Todoinfo, error){
	return  repository.Gettodolist()
}