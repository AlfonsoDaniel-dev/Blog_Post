package main

import (
	"github.com/TeenBanner/Inventory_system/cmd/config"
	"github.com/TeenBanner/Inventory_system/pkg/authorization"
	"github.com/TeenBanner/Inventory_system/pkg/database"
	"github.com/TeenBanner/Inventory_system/pkg/database/migrations"
	"log"
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

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Can't start server: %v", err)
	}
}
