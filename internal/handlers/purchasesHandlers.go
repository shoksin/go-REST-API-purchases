package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/services"
	"github.com/shoksin/go-REST-API-purchases/middleware"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PurchasesHandler interface {
	CreatePurchase(c echo.Context) error
	GetPurchases(c echo.Context) error
	DeletePurchase(c echo.Context) error
	DeleteUserPurchases(c echo.Context) error
	UpdateUserPurchase(c echo.Context) error
}

type purchasesHandler struct {
	purchasesService services.PurchasesService
	logger           logging.Logger
}

func NewPurchasesHandler(purchasesService services.PurchasesService, logger logging.Logger) PurchasesHandler {
	return &purchasesHandler{purchasesService: purchasesService, logger: logger.GetLoggerWithField("layer", "PurchasesHandlers")}
}

// CreatePurchase @Summary Создание покупки
// @Tags purchases
// @Description Создание новой покупки
// @ID create
// @Accept  json
// @Produce  json
// @Param input body swagger.CreatePurchase true "Данные о покупке"
// @Success 201 {object} swagger.CreatePurchaseResponse
// @Security ApiKeyAuth
// @Router /purchases/ [post]
func (p *purchasesHandler) CreatePurchase(c echo.Context) error {
	purchase := &models.Purchase{}
	if err := c.Bind(&purchase); err != nil {
		return utils.Respond(c, http.StatusBadRequest, utils.Message("not right request body"))
	}

	tk, _ := middleware.GetToken(c)

	purchase.UserID = tk.UserId
	if validResp := purchase.Validate(); validResp != nil {
		p.logger.WithFields(logrus.Fields{
			"userID":   tk.UserId,
			"name":     purchase.Name,
			"price":    purchase.Price,
			"quantity": purchase.Quantity,
		}).Warn("invalid purchase data")
		return utils.Respond(c, http.StatusBadRequest, validResp)
	}
	resp, err := p.purchasesService.CreatePurchase(purchase)
	if err != nil {
		return utils.Respond(c, http.StatusBadRequest, resp)
	}
	return utils.Respond(c, http.StatusCreated, resp)
}

// GetPurchases @Summary Получить покупки пользователя
// @Tags purchases
// @Description Получение покупок пользователя по id
// @ID get
// @Accept  json
// @Produce  json
// @Success 201 {object} swagger.GetPurchasesResponse
// @Security ApiKeyAuth
// @Router /purchases/ [Get]
func (p *purchasesHandler) GetPurchases(c echo.Context) error {
	tk, _ := middleware.GetToken(c)
	purchases, err := p.purchasesService.GetPurchases(tk.UserId)
	if err != nil {
		return utils.Respond(c, http.StatusInternalServerError, purchases)
	}
	return utils.Respond(c, http.StatusOK, purchases)
}

// DeletePurchase @Summary Удаление покупки
// @Tags purchases
// @Description Удаление конкретной покупки по ID
// @ID delete-purchase
// @Produce json
// @Param id path int true "ID покупки"
// @Success 200 {object} swagger.DeletePurchaseResponse
// @Failure 403 {object} swagger.DeletePurchaseResponse
// @Failure 404 {object} swagger.DeletePurchaseResponse
// @Security ApiKeyAuth
// @Router /purchases/{id} [delete]
func (p *purchasesHandler) DeletePurchase(c echo.Context) error {
	purchaseId := c.Param("id")
	id, ok := strconv.Atoi(purchaseId)
	if ok != nil || id < 0 {
		p.logger.WithFields(logrus.Fields{
			"purchase_id": purchaseId,
		}).Warn("invalid purchase_id")
		return utils.Respond(c, http.StatusBadRequest, utils.Message("purchase id must be positive integer"))
	}
	resp, err := p.purchasesService.DeletePurchase(uint(id))
	if err != nil {
		return utils.Respond(c, http.StatusNotFound, resp)
	}
	return utils.Respond(c, http.StatusOK, resp)
}

// DeleteUserPurchases @Summary Удаление всех покупок пользователя
// @Tags purchases
// @Description Удаление всех покупок пользователя
// @ID delete-purchases
// @Produce json
// @Success 200 {object} swagger.DeletePurchaseResponse
// @Failure 403 {object} swagger.DeletePurchaseResponse
// @Failure 404 {object} swagger.DeletePurchaseResponse
// @Security ApiKeyAuth
// @Router /purchases/ [delete]
func (p *purchasesHandler) DeleteUserPurchases(c echo.Context) error {
	tk, _ := middleware.GetToken(c)
	resp, err := p.purchasesService.DeleteUserPurchases(tk.UserId)
	if err != nil {
		return utils.Respond(c, http.StatusNotFound, resp)
	}
	return utils.Respond(c, http.StatusOK, resp)
}

func (p *purchasesHandler) UpdateUserPurchase(c echo.Context) error {
	purchaseId := c.QueryParam("id")
	id, ok := strconv.Atoi(purchaseId)
	if ok != nil || id < 0 {
		p.logger.WithFields(logrus.Fields{
			"purchase_id": purchaseId,
		}).Warn("invalid purchase_id")
		return utils.Respond(c, http.StatusBadRequest, utils.Message("purchase id must be positive integer"))
	}
	tk, _ := middleware.GetToken(c)
	purchase := &models.Purchase{}
	if err := c.Bind(&purchase); err != nil {
		return utils.Respond(c, http.StatusBadRequest, utils.Message("not right request body"))
	}

	purchase.UserID = tk.UserId

	if err := purchase.Validate(); err != nil {
		p.logger.Warn("invalid purchase data")
	}
	resp, err := p.purchasesService.UpdateUserPurchase(uint(id), purchase)
	if err != nil {
		return utils.Respond(c, http.StatusInternalServerError, resp)
	}
	return utils.Respond(c, http.StatusOK, resp)
}
