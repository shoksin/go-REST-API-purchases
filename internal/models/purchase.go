package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int64   `json:"quantity"`
	FullPrice float64 `json:"full_price"`
}

func (p *Purchase) calculateFullPrice() {
	p.FullPrice = float64(p.Quantity) * p.Price
}
