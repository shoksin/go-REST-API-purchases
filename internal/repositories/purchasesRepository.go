package repositories

import (
	"errors"
	"github.com/shoksin/go-REST-API-purchases/internal/db"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PurchasesRepository interface {
	Create(purchase *models.Purchase) (*models.Purchase, error)
	GetPurchases(userId uint) ([]*models.Purchase, error)
	DeletePurchase(id uint) error
	DeletePurchases(userId uint) error
	UpdatePurchase(purchaseId uint, purchase *models.Purchase) (*models.Purchase, error)
}

type purchasesRepository struct {
	db     *gorm.DB
	logger logging.Logger
}

func NewPurchasesRepository(db *gorm.DB, logger logging.Logger) PurchasesRepository {
	return &purchasesRepository{db, logger.GetLoggerWithField("layer", "PurchasesRepository")}
}

func (p *purchasesRepository) Create(purchase *models.Purchase) (*models.Purchase, error) {
	purchase.CalculateFullPrice()
	purchase.FullPrice = float64(purchase.FullPrice)
	if err := db.GetDB().Create(purchase).Error; err != nil {
		p.logger.WithFields(
			logrus.Fields{
				"purchase name": purchase.Name,
				"user_id":       purchase.UserID,
			}).Error("Failed to create purchase")
		return nil, err
	}
	return purchase, nil
}

func (p *purchasesRepository) GetPurchases(userId uint) ([]*models.Purchase, error) {
	purchases := make([]*models.Purchase, 0)
	if err := db.GetDB().Table("purchases").Where("user_id = ?", userId).Find(&purchases).Error; err != nil {
		p.logger.WithField("user_id", userId).Error("Failed to get purchase")
		return nil, err
	}
	return purchases, nil
}

func (p *purchasesRepository) DeletePurchase(id uint) error {
	result := db.GetDB().Table("purchases").Where("id = ?", id).Delete(&models.Purchase{})
	if result.Error != nil {
		p.logger.WithField("purchase_id", id).Error("Failed to delete purchase")
		return result.Error
	}
	if result.RowsAffected == 0 {
		p.logger.WithField("purchase_id", id).Warning("Purchase not found for deletion")
		return errors.New("purchase not found")
	}
	return nil
}

func (p *purchasesRepository) DeletePurchases(userId uint) error {
	result := db.GetDB().Table("purchases").Where("user_id = ?", userId).Delete(&models.Purchase{})
	if result.Error != nil {
		p.logger.WithField("user_id", userId).Error("Failed to delete purchases for user")
		return result.Error
	}
	if result.RowsAffected == 0 {
		p.logger.WithField("user_id", userId).Warning("No purchases found for user deletion")
		return errors.New("purchase not found")
	}
	p.logger.WithFields(logrus.Fields{
		"user_id":       userId,
		"rows_affected": result.RowsAffected,
	}).Info("Deleted purchases for user")
	return nil
}

func (p *purchasesRepository) UpdatePurchase(purchaseId uint, purchase *models.Purchase) (*models.Purchase, error) {
	purchase.FullPrice = float64(purchase.FullPrice)
	tempPurchase := &models.Purchase{}
	err := db.GetDB().Table("purchases").Where("id = ?", purchaseId).First(tempPurchase).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			p.logger.Info("purchase not found")
			return nil, errors.New("purchase not found")
		}
		p.logger.Error("Error while fetching purchase data from DB")
		return nil, err
	}

	tempPurchase.Assign(purchase)

	if err := db.GetDB().Save(&tempPurchase).Error; err != nil {
		p.logger.WithFields(
			logrus.Fields{
				"purchase name": purchase.Name,
				"user_id":       purchase.UserID,
			}).Error("Failed to update purchase")
	}
	return purchase, nil
}
