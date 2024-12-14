package services

import (
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-REST-API-purchases/internal/repositories"
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
)

type PurchasesService interface {
	CreatePurchase(purchase *models.Purchase) (map[string]interface{}, error)
	GetPurchases(userId uint) (map[string]interface{}, error)
	DeletePurchase(id uint) (map[string]interface{}, error)
	DeleteUserPurchases(userId uint) (map[string]interface{}, error)
}

type purchasesService struct {
	purchaseRepository repositories.PurchasesRepository
	logger             logging.Logger
}

func NewPurchasesService(purchaseRepository repositories.PurchasesRepository, logger logging.Logger) PurchasesService {
	return &purchasesService{purchaseRepository: purchaseRepository, logger: logger.GetLoggerWithField("layer", "PurchasesService")}
}

func (p *purchasesService) CreatePurchase(purchase *models.Purchase) (map[string]interface{}, error) {
	purchaseResp, err := p.purchaseRepository.Create(purchase)
	if err != nil {
		return utils.Message("creation failed"), err
	}
	resp := utils.Message("purchase created")
	resp["purchase"] = purchaseResp
	return resp, nil
}

func (p *purchasesService) GetPurchases(userId uint) (map[string]interface{}, error) {
	purchasesResp, err := p.purchaseRepository.GetPurchases(userId)
	if err != nil {
		return utils.Message("purchase not found"), err
	}
	resp := utils.Message("purchase found")
	if len(purchasesResp) == 0 {
		resp = utils.Message("There is no purchases")
	}
	resp["purchases"] = purchasesResp
	return resp, nil
}

func (p *purchasesService) DeletePurchase(id uint) (map[string]interface{}, error) {
	err := p.purchaseRepository.DeletePurchase(id)
	if err != nil {
		return utils.Message("purchase not found"), err
	}
	resp := utils.Message("purchase deleted")
	return resp, nil
}

func (p *purchasesService) DeleteUserPurchases(userId uint) (map[string]interface{}, error) {
	err := p.purchaseRepository.DeletePurchases(userId)
	if err != nil {
		return utils.Message("purchase not found"), err
	}
	resp := utils.Message("purchases deleted")
	return resp, nil
}
