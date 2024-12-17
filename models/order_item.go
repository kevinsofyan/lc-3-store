package models

type OrderItem struct {
	OrderItemID int     `json:"order_item_id" gorm:"primaryKey"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
