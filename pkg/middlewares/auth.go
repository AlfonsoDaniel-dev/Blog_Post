package middlewares

import (
	"github.com/TeenBanner/Inventory_system/User/infrastructure/controllers/responses"
	"github.com/TeenBanner/Inventory_system/pkg/authorization"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			response := responses.NewResponse(err, "Error", "Not authorized")
			return c.JSON(http.StatusForbidden, response)
		}

		return f(c)
	}
}
