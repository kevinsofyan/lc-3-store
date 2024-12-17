package repositories

import (
	"store/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllCarts() ([]models.Cart, error)
	AddCart(cart *models.Cart) error
	DeleteCart(id string) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) GetAllCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Find(&carts).Error
	return carts, err
}

func (r *cartRepository) AddCart(cart *models.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) DeleteCart(id string) error {
	return r.db.Delete(&models.Cart{}, id).Error
}
