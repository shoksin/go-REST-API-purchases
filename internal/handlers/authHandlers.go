package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/services"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, utils.Message(false, "bad request"))
	}

	resp := h.userService.Create(user)
	return c.JSON(http.StatusCreated, resp)
}
