package config

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/pkg/API"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHttp(db *sql.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger(), middleware.Recover())

	API.InitRoutes(e, db)

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5500", "*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}

	e.Use(middleware.CORSWithConfig(corsConfig))

	return e
}
