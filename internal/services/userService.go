package services

import (
	"fmt"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/repositories"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
)

type UserService interface {
	Create(user *models.User) (map[string]interface{}, error)
	Login(email, password string) (map[string]interface{}, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Create(user *models.User) (map[string]interface{}, error) {
	userResp, err := s.userRepository.CreateUser(user)
	if err != nil {
		logging.GetLogger().Error(err)
		return utils.Message("Register failed"), err
	}
	logging.GetLogger().Debug(fmt.Println(user))
	resp := utils.Message("Account created!")
	resp["user"] = userResp
	return resp, nil
}

func (s *userService) Login(email, password string) (map[string]interface{}, error) {
	userResp, err := s.userRepository.Login(email, password)
	if err != nil {
		logging.GetLogger().Error(err)
		return utils.Message(err.Error()), err
	}
	resp := utils.Message("Login successful!")
	resp["user"] = userResp
	return resp, nil
}
