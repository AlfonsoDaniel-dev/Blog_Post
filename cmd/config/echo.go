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

	/*
		corsConfig := middleware.CORSConfig{
			AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
			AllowMethods: strings.Split(os.Getenv("ALLOWED_METHODS"), ","),
		}

		e.Use(middleware.CORSWithConfig(corsConfig)) */

	return e
}
