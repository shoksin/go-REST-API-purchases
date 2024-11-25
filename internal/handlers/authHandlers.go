package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/services"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"net/http"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
	Login(c echo.Context) error
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return utils.Respond(c, http.StatusBadRequest, utils.Message("Bad request"))
	}

	if validResp := user.ValidateRegister(); validResp != nil {
		return utils.Respond(c, http.StatusBadRequest, validResp)
	}
	resp, err := h.userService.Create(user)
	if err != nil {

	}
	return utils.Respond(c, http.StatusCreated, resp)
}

func (h *userHandler) Login(c echo.Context) error {
	loginUser := &models.LoginUser{}
	if err := c.Bind(loginUser); err != nil {
		return utils.Respond(c, http.StatusBadRequest, utils.Message("Bad request"))
	}
	if validResp := loginUser.ValidateLogin(); validResp != nil {
		return utils.Respond(c, http.StatusBadRequest, validResp)
	}
	resp, err := h.userService.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		return utils.Respond(c, http.StatusForbidden, resp)
	}
	return utils.Respond(c, http.StatusOK, resp)
}
