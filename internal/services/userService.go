package services

import (
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/repositories"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	Create(user *models.User) (map[string]interface{}, error)
	Login(email, password string) (map[string]interface{}, error)
}

type userService struct {
	userRepository repositories.UserRepository
	logger         logging.Logger
}

func NewUserService(userRepository repositories.UserRepository, logger logging.Logger) UserService {
	return &userService{userRepository: userRepository, logger: logger.GetLoggerWithField("layer", "UserService")}
}

func (s *userService) Create(user *models.User) (map[string]interface{}, error) {
	isTaken, err := s.userRepository.IsEmailTaken(user.Email)
	if err != nil {
		return utils.Message("Error checking email existence"), err
	}

	if isTaken {
		s.logger.WithField("email", user.Email).Info("Email already registered")
		return utils.Message("Email already taken"), nil
	}
	userResp, err := s.userRepository.CreateUser(user)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"email": user.Email,
			"name":  user.Name,
		}).Error("Error while creating user")
		return utils.Message("Register failed"), err
	}
	resp := utils.Message("Account created!")
	resp["user"] = userResp
	return resp, nil
}

func (s *userService) Login(email, password string) (map[string]interface{}, error) {
	userResp, err := s.userRepository.Login(email, password)
	if err != nil {
		s.logger.Error("Failed to login")
		return utils.Message(err.Error()), err
	}
	resp := utils.Message("Login successful!")
	resp["user"] = userResp
	return resp, nil
}
