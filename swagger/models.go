package swagger

import (
	"time"
)

type UserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type User struct {
	GormModel
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"date_of_birth" time_format:"2006-01-02"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password"`
	Role        string    `json:"role" default:"user"`
}

type RegisterUser struct {
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"date_of_birth" time_format:"2006-01-02"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password"`
	Role        string    `json:"role" default:"user"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" default:"user"`
}

type LoginResponse struct {
	Message string `json:"message"`
	GormModel
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" default:"user"`
	Token    string `json:"token" sql:"-"`
}

type Purchase struct {
	GormModel
	UserID    uint    `json:"user_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
	FullPrice float64 `json:"full_price"`
	Comment   string  `json:"comment"`
}

type CreatePurchase struct {
	UserID   uint    `json:"user_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity uint64  `json:"quantity"`
	Comment  string  `json:"comment"`
}

type CreatePurchaseResponse struct {
	Message string `json:"message"`
	GormModel
	UserID    uint    `json:"user_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
	FullPrice float64 `json:"full_price"`
	Comment   string  `json:"comment"`
}

type GetPurchasesResponse struct {
	Message   string     `json:"message"`
	Purchases []Purchase `json:"purchases"`
}

type DeletePurchaseResponse struct {
	Message string `json:"message"`
}

type GormModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"CreatedAt" example:"2024-02-03T15:04:05Z"`
	UpdatedAt time.Time `json:"UpdatedAt" example:"2024-02-03T15:04:05Z"`
	DeletedAt time.Time `gorm:"index" json:"DeletedAt"`
}
