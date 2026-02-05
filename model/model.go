package main

import "gorm.io/gorm"

type Todo_info struct{
	gorm.Model
	TodoId uint `json:"todoid" gorm:"primary key"`
	Title string `json:"title"`
}
type User struct{
	gorm.Model
	ID uint `json:"id" gorm:"primary key"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}