package models

import "time"

type Token struct {
	TokenID   int       `json:"token_id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	JWTToken  string    `json:"jwt_token"`
	CreatedAt time.Time `json:"created_at"`
}
