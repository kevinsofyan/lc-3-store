package routes

import (
	"store/config"
	"store/controllers"
	"store/middlewares"
	"store/repositories"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository(config.DB)
	tokenRepo := repositories.NewTokenRepository(config.DB)
	cartRepo := repositories.NewCartRepository(config.DB)
	orderRepo := repositories.NewOrderRepository(config.DB)
	productRepo := repositories.NewProductRepository(config.DB)

	userController := controllers.NewUserController(userRepo, tokenRepo)
	cartController := controllers.NewCartController(cartRepo)
	orderController := controllers.NewOrderController(orderRepo)
	productController := controllers.NewProductController(productRepo)

	userGroup := e.Group("/users")
	userGroup.POST("/register", userController.RegisterUser)
	userGroup.POST("/login", userController.LoginUser)

	cartGroup := userGroup.Group("/carts", middlewares.JWTMiddleware)
	cartGroup.GET("", cartController.GetCarts)
	cartGroup.POST("", cartController.AddCart)
	cartGroup.DELETE("/:id", cartController.DeleteCart)

	orderGroup := userGroup.Group("/orders", middlewares.JWTMiddleware)
	orderGroup.GET("", orderController.GetOrders)
	orderGroup.POST("", orderController.CreateOrder)

	productGroup := e.Group("/products")
	productGroup.GET("", productController.GetProducts)
	productGroup.GET("/:id", productController.GetProductByID)
}
