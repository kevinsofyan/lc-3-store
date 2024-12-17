package repositories

import (
	"store/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
	CreateOrder(userID int) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) CreateOrder(userID int) error {
	var carts []models.Cart
	if err := r.db.Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return err
	}

	var totalPrice float64
	for _, cart := range carts {
		var product models.Product
		if err := r.db.Where("product_id = ?", cart.ProductID).First(&product).Error; err != nil {
			return err
		}
		totalPrice += product.Price * float64(cart.Quantity)
	}

	order := models.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
	}
	if err := r.db.Create(&order).Error; err != nil {
		return err
	}

	for _, cart := range carts {
		var product models.Product
		if err := r.db.Where("product_id = ?", cart.ProductID).First(&product).Error; err != nil {
			return err
		}
		orderItem := models.OrderItem{
			OrderID:   order.OrderID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     product.Price * float64(cart.Quantity),
		}
		if err := r.db.Create(&orderItem).Error; err != nil {
			return err
		}
	}

	if err := r.db.Where("user_id = ?", userID).Delete(&models.Cart{}).Error; err != nil {
		return err
	}

	return nil
}
