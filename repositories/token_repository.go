package repositories

import (
	"store/models"

	"gorm.io/gorm"
)

type TokenRepository interface {
	SaveToken(token *models.Token) error
	FindTokenByUserID(userID int) (*models.Token, error)
	UpdateToken(token *models.Token) error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db}
}

func (r *tokenRepository) SaveToken(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r *tokenRepository) FindTokenByUserID(userID int) (*models.Token, error) {
	var token models.Token
	err := r.db.Where("user_id = ?", userID).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *tokenRepository) UpdateToken(token *models.Token) error {
	return r.db.Save(token).Error
}
