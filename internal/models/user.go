package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"gorm.io/gorm"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int64  `json:"age"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

type LoginUser struct {
	gorm.Model
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token" sql:"-"`
}

func (u User) ValidateRegister() map[string]interface{} {
	if !strings.Contains(u.Email, "@") {
		return utils.Message("Email should contain '@'")
	}
	if len(u.Password) < 8 {
		return utils.Message("Password should contain at least 8 characters")
	}
	return nil
}

func (u LoginUser) ValidateLogin() map[string]interface{} {
	if !strings.Contains(u.Email, "@") {
		return utils.Message("Email should contain '@'")
	}
	if len(u.Password) < 8 {
		return utils.Message("Password must be at least 8 characters")
	}
	return nil
}
