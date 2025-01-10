package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/services"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
	Login(c echo.Context) error
}

type userHandler struct {
	userService services.UserService
	logger      logging.Logger
}

func NewUserHandler(userService services.UserService, logger logging.Logger) UserHandler {
	return &userHandler{userService, logger.GetLoggerWithField("layer", "AuthHandlers")}
}

// CreateUser godoc
// @Summary CreateUser
// @Tags auth
// @Summary create account
// @ID signup
// @Accept json
// @Produce json
// @Param input body models.User true "account info"
func (h *userHandler) CreateUser(c echo.Context) error {
	user := &models.User{}
	mockUser := &models.MockUser{}
	if err := c.Bind(mockUser); err != nil {
		h.logger.WithFields(logrus.Fields{
			"request body": c.Request().Body,
		}).Error("Unable to bind request body")
		return utils.Respond(c, http.StatusBadRequest, utils.Message("Bad request"))
	}

	user.Name = mockUser.Name
	user.Surname = mockUser.Surname
	user.Email = mockUser.Email
	user.Password = mockUser.Password
	user.Role = mockUser.Role

	t, err := time.Parse("2006-01-02", mockUser.DateOfBirth)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
			"date_of_birth": mockUser.DateOfBirth,
		})
	}

	user.DateOfBirth = t

	//if err := c.Bind(user); err != nil {
	//	h.logger.WithFields(logrus.Fields{
	//		"request body": c.Request().Body,
	//	}).Error("Unable to bind request body")
	//	return utils.Respond(c, http.StatusBadRequest, utils.Message("Bad request"))
	//}

	if validResp := user.ValidateRegister(); validResp != nil {
		h.logger.WithFields(logrus.Fields{
			"email": user.Email,
			"name":  user.Name,
		}).Warning("Not valid response")
		return utils.Respond(c, http.StatusBadRequest, validResp)
	}
	resp, err := h.userService.Create(user)
	if err != nil {
		return utils.Respond(c, http.StatusBadRequest, resp)
	}
	return utils.Respond(c, http.StatusCreated, resp)
}

func (h *userHandler) Login(c echo.Context) error {
	loginUser := &models.LoginUser{}
	if err := c.Bind(loginUser); err != nil {
		return utils.Respond(c, http.StatusBadRequest, utils.Message("Bad request"))
	}
	if validResp := loginUser.ValidateLogin(); validResp != nil {
		h.logger.WithFields(logrus.Fields{
			"email": loginUser.Email,
		}).Warning("Not valid response")
		return utils.Respond(c, http.StatusBadRequest, validResp)
	}
	resp, err := h.userService.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		return utils.Respond(c, http.StatusForbidden, resp)
	}
	return utils.Respond(c, http.StatusOK, resp)
}
