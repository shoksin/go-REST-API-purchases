package models

import (
	"github.com/shoksin/go-REST-API-purchases/pkg/utils"
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
	FullPrice float64 `json:"full_price"`
	Comment   string  `json:"comment"`
}

func (p *Purchase) CalculateFullPrice() {
	p.FullPrice = float64(p.Quantity) * p.Price
}

func (p *Purchase) Validate() map[string]interface{} {
	if p.Quantity == 0 {
		return utils.Message("Quantity shouldn't be zero")
	}
	if p.Price < 0 {
		return utils.Message("Price shouldn't be negative")
	}
	if p.Name == "" {
		return utils.Message("Name shouldn't be empty")
	}
	return nil
}

func (p *Purchase) Assign(p2 *Purchase) {
	if p2.Name != "" {
		p.Name = p2.Name
	}
	if p2.Price > 0 {
		p.Price = p2.Price
	}
	if p2.Quantity > 0 {
		p.Quantity = p2.Quantity
	}
	if p2.Comment != "" {
		p.Comment = p2.Comment
	}
	p.CalculateFullPrice()
}
