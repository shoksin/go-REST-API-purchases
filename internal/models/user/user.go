package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int64  `json:"age"`
	Email   string `json:"email"`
}
