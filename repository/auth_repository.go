package repository

import (
	database "todo_api_backend/db"
	"todo_api_backend/model"
)

func CreateUser(user *model.User)error{
  db := database.Databaseconnection()
  return db.Create(&user).Error
}
func Userlogin(username string)(*model.User, error){
	db := database.Databaseconnection()
	var user model.User
	
	 if err := db.Where("username = ?", username).First(&user).Error; err != nil{
		return  nil , err
	 }
	 
	return  &user, nil

	
}