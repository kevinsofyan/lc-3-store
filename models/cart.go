package models

import "time"

type Cart struct {
	CartID    int       `json:"cart_id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
