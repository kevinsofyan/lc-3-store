package controllers

import (
	"net/http"
	"store/repositories"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	repo repositories.OrderRepository
}

func NewOrderController(repo repositories.OrderRepository) *OrderController {
	return &OrderController{repo}
}

// GetOrders godoc
// @Summary menampilkan seluruh data order
// @Description menampilkan seluruh data order
// @Tags order
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Produce json
// @Success 200 {array} models.Order
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/orders [get]
func (ctrl *OrderController) GetOrders(c echo.Context) error {
	orders, err := ctrl.repo.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to retrieve orders"})
	}
	return c.JSON(http.StatusOK, orders)
}

// CreateOrder godoc
// @Summary membuat data order
// @Description membuat data order baru berdasarkan data product yang tersimpan pada carts
// @Tags order
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/orders [post]
func (ctrl *OrderController) CreateOrder(c echo.Context) error {
	userID := c.Get("userID").(int)

	if err := ctrl.repo.CreateOrder(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create order"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create order data",
	})
}
