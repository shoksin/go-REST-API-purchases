package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

type Token struct {
	UserId uint
	Role   string
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"date_of_birth" time_format:"2006-01-02"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password"`
	Role        string    `json:"role" default:"user"`
}

type MockUser struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DateOfBirth string `json:"date_of_birth" time_format:"2006-01-02"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password"`
	Role        string `json:"role" default:"user"`
}

type LoginUser struct {
	gorm.Model
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" default:"user"`
	Token    string `json:"token" sql:"-"`
}

func (u User) ValidateRegister() map[string]interface{} {
	if !strings.Contains(u.Email, "@") {
		return utils.Message("Email should contain '@'")
	}
	if len(u.Password) < 8 {
		return utils.Message("Password should contain at least 8 characters")
	}
	if u.DateOfBirth.After(time.Now()) {
		return utils.Message("Date of birth cannot be in the future")
	}
	if u.Role != "user" && u.Role != "admin" {
		return utils.Message("Invalid role")
	}

	if u.Role == "admin" && u.Password != os.Getenv("ADMIN_PASSWORD") {
		return utils.Message("Invalid admin password")
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
