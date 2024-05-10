package main

import (
	"github.com/TeenBanner/Inventory_system/cmd/config"
	"github.com/TeenBanner/Inventory_system/database"
	"github.com/TeenBanner/Inventory_system/database/migrations"
	"log"
)

func main() {
	err := config.LoadEnv()

	if err := config.ValidateEnvVars(); err != nil {
		log.Fatalf("Cannot validate .env vars")
	}

	if err != nil {
		log.Fatalf("Can't load .env")
	}
	connStr := config.CreateStrConn()

	database.CreateConnection(connStr)

	db := database.Pool()

	migrator := migrations.NewMigrator(db)

	if err := migrator.Migrate(db); err != nil {
		log.Fatalf("no se pudieron hacer las migraciones. Error: %v", err)
	}

}
