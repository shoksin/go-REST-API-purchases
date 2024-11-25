package repositories

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/shoksin/go-REST-API-purchases/internal/db"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	Login(email string, password string) (*models.LoginUser, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) CreateUser(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	if err := db.GetDB().Create(user).Error; err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func (repo *userRepository) Login(email string, password string) (*models.LoginUser, error) {
	account := &models.LoginUser{}
	if err := db.GetDB().Table("users").Where("email = ?", email).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user with this email does not exist")
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		logging.GetLogger().Debug("Not correct password")
		return nil, errors.New("wrong password")
	}

	account.Password = ""

	tk := &models.Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	account.Token = tokenString
	return account, nil
}
