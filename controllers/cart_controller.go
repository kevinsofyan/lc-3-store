package controllers

import (
	"net/http"
	"store/models"
	"store/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CartController struct {
	repo repositories.CartRepository
}

func NewCartController(repo repositories.CartRepository) *CartController {
	return &CartController{repo}
}

// GetCarts godoc
// @Summary menampilkan seluruh data cart
// @Description menampilkan seluruh data cart
// @Tags cart
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Produce json
// @Success 200 {array} models.Cart
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/carts [get]
func (ctrl *CartController) GetCarts(c echo.Context) error {
	carts, err := ctrl.repo.GetAllCarts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to retrieve carts"})
	}
	return c.JSON(http.StatusOK, carts)
}

// AddCart godoc
// @Summary menambahkan data cart
// @Description menambahkan data cart
// @Tags cart
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Param cart body models.Cart true "Cart"
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/carts [post]
func (ctrl *CartController) AddCart(c echo.Context) error {
	cart := new(models.Cart)
	if err := c.Bind(cart); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input"})
	}

	if err := ctrl.repo.AddCart(cart); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to add cart"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add product to cart",
	})
}

func (ctrl *CartController) DeleteCart(c echo.Context) error {
	id := c.Param("id")
	if err := ctrl.repo.DeleteCart(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{Message: "Cart not found"})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to delete cart"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success remove product from cart",
	})
}
