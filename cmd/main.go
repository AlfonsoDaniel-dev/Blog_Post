package main

import (
	"fmt"
	"github.com/TeenBanner/Inventory_system/cmd/config"
	"github.com/TeenBanner/Inventory_system/pkg/authorization"
	"github.com/TeenBanner/Inventory_system/pkg/database"
	"github.com/TeenBanner/Inventory_system/pkg/database/migrations"
	"log"
	"os"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal(err)
	}
	if err := config.ValidateEnvVars(); err != nil {
		log.Fatalf("Cannot validate .env vars")
	}

	if err := authorization.LoadFile("certificates/app.rsa", "certificates/app.rsa.pub"); err != nil {
		log.Fatalf("No se pudieron cargar los certificados. ERR: %v", err)
	}

	connStr := config.CreateStrConn()

	database.CreateConnection(connStr)

	db := database.Pool()

	migrator := migrations.NewMigrator(db)

	if err := migrator.Migrate(db); err != nil {
		log.Fatalf("no se pudieron hacer las migraciones. Error: %v", err)
	}

	e := config.NewHttp(db)

	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")

	dir := fmt.Sprintf("%v:%v", host, port)

	if err := e.Start(dir); err != nil {
		log.Fatalf("Can't start server: %v", err)
	}
}
