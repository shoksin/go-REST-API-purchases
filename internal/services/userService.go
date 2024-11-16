package services

import (
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/repositories"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
)

type UserService interface {
	Create(user *models.User) map[string]interface{}
	Login(email, password string) map[string]interface{}
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Create(user *models.User) map[string]interface{} {
	err := s.userRepository.CreateUser(user)
	if err != nil {
		logging.GetLogger().Fatal("Account wasn't created!")
		return utils.Message(false, "Register failed")
	}
	return utils.Message(true, "Account created!")
}

func (s *userService) Login(email, password string) map[string]interface{} {
	err := s.userRepository.Login(email, password)
	if err != nil {
		logging.GetLogger().Fatal("Login failed!")
		return utils.Message(false, "Login failed!")
	}
	return utils.Message(true, "Login successful!")
}
