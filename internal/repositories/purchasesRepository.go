package repositories

import (
	"errors"
	"github.com/shoksin/go-REST-API-purchases/internal/db"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"gorm.io/gorm"
)

type PurchasesRepository interface {
	Create(purchase *models.Purchase) (*models.Purchase, error)
	GetPurchases(userId uint) ([]*models.Purchase, error)
	DeletePurchase(id uint) error
	DeletePurchases(userId uint) error
}

type purchasesRepository struct {
	db *gorm.DB
}

func NewPurchasesRepository(db *gorm.DB) PurchasesRepository {
	return &purchasesRepository{db}
}

func (p *purchasesRepository) Create(purchase *models.Purchase) (*models.Purchase, error) {
	purchase.CalculateFullPrice()
	purchase.FullPrice = float64(purchase.FullPrice)
	if err := db.GetDB().Create(purchase).Error; err != nil {
		return nil, err
	}
	return purchase, nil
}

func (p *purchasesRepository) GetPurchases(userId uint) ([]*models.Purchase, error) {
	purchases := make([]*models.Purchase, 0)
	if err := db.GetDB().Table("purchases").Where("user_id = ?", userId).Find(&purchases).Error; err != nil {
		return nil, err
	}
	return purchases, nil
}

func (p *purchasesRepository) DeletePurchase(id uint) error {
	result := db.GetDB().Table("purchases").Where("id = ?", id).Delete(&models.Purchase{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("purchase not found")
	}
	return nil
}

func (p *purchasesRepository) DeletePurchases(userId uint) error {
	result := db.GetDB().Table("purchases").Where("user_id = ?", userId).Delete(&models.Purchase{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("purchase not found")
	}
	return nil
}
