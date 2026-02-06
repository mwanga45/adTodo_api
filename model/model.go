package model

import "gorm.io/gorm"

type Todoinfo struct{
	gorm.Model
	Title string `json:"title"`

}
type User struct{
	gorm.Model
	Username  string `json:"username"`
	Complete bool `json:"complete"`
	Password  string `json:"password"`

}
