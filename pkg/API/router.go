package API

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/User/infrastructure/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *sql.DB) {
	controller := controllers.BuildUserController(e, db)

	controller.PublicRoutes(e)
	controller.PrivateRoutes(e)
}
