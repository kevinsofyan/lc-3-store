package models

import "time"

type Order struct {
	OrderID    int       `json:"order_id" gorm:"primaryKey"`
	UserID     int       `json:"user_id"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}
