package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/services"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"github.com/sirupsen/logrus"
	"net/http"
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

func (h *userHandler) CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		h.logger.WithFields(logrus.Fields{
			"request body": c.Request().Body,
		}).Error("Unable to bind request body")
		return utils.Respond(c, http.StatusBadRequest, utils.Message("Bad request"))
	}

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
