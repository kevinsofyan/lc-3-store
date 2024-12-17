package middlewares

import (
	"net/http"
	"store/models"
	"store/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Unauthorized"})
		}

		// Split the "Bearer" part and the token part
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Unauthorized"})
		}

		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Unauthorized"})
		}

		// Store user ID in context
		c.Set("userID", claims.UserID)

		return next(c)
	}
}
