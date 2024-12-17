package models

type User struct {
	UserID   int    `json:"user_id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSuccess struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
