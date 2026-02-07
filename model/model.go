package model

import "gorm.io/gorm"

type Todoinfo struct{
	gorm.Model
	Title string `json:"title"`
	Confirm bool `json:"confirm"`

}
type User struct{
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`

}
