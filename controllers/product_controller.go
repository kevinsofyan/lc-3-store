package controllers

import (
	"net/http"
	"store/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductController struct {
	repo repositories.ProductRepository
}

func NewProductController(repo repositories.ProductRepository) *ProductController {
	return &ProductController{repo}
}

// GetProducts godoc
// @Summary menampilkan semua data product yang tersedia
// @Description menampilkan semua data product yang tersedia
// @Tags product
// @Produce json
// @Success 200 {array} models.Product
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products [get]
func (ctrl *ProductController) GetProducts(c echo.Context) error {
	products, err := ctrl.repo.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to retrieve products"})
	}
	return c.JSON(http.StatusOK, products)
}

// GetProductByID godoc
// @Summary menampilkan data product sesuai dengan id
// @Description menampilkan semua data product sesuai dengan id
// @Tags product
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} models.Product
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products/{id} [get]
func (ctrl *ProductController) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	product, err := ctrl.repo.GetProductByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{Message: "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to retrieve product"})
	}
	return c.JSON(http.StatusOK, product)
}
