package controllers

import (
	"net/http"
	"store/models"
	"store/repositories"
	"store/utils"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	repo      repositories.UserRepository
	tokenRepo repositories.TokenRepository
}

func NewUserController(repo repositories.UserRepository, tokenRepo repositories.TokenRepository) *UserController {
	return &UserController{repo, tokenRepo}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// RegisterUser godoc
// @Summary register user baru
// @Description register user baru ke dalam aplikasi
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/register [post]
func (ctrl *UserController) RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input"})
	}

	existingUser, err := ctrl.repo.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Email already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to hash password"})
	}

	user.Password = string(hashedPassword)

	if err := ctrl.repo.Register(user); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to register user"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success register",
		"user": map[string]interface{}{
			"user_id": user.UserID,
			"name":    user.Name,
			"email":   user.Email,
		},
	})
}

// LoginUser godoc
// @Summary login user yang sudah terdaftar
// @Description login user yang sebelumnya sudah berhasil melakukan registrasi, akan membalikan data akses token
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.LoginRequest true "Login Request"
// @Success 200 {object} models.LoginSuccess
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/login [post]
func (ctrl *UserController) LoginUser(c echo.Context) error {
	loginRequest := new(models.LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input"})
	}

	storedUser, err := ctrl.repo.FindByEmail(loginRequest.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid email or password"})
	}

	if bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginRequest.Password)) != nil {
		if storedUser.Password != loginRequest.Password {
			return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid email or password"})
		}
	}

	token, err := utils.GenerateJWT(uint(storedUser.UserID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to generate token"})
	}

	existingToken, err := ctrl.tokenRepo.FindTokenByUserID(storedUser.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to find token"})
	}

	if existingToken != nil {
		existingToken.JWTToken = token
		existingToken.CreatedAt = time.Now()
		if err := ctrl.tokenRepo.UpdateToken(existingToken); err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to update token"})
		}
	} else {
		// Save the new token
		tokenModel := &models.Token{
			UserID:    storedUser.UserID,
			JWTToken:  token,
			CreatedAt: time.Now(),
		}
		if err := ctrl.tokenRepo.SaveToken(tokenModel); err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to save token"})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}
