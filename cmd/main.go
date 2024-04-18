package main

import (
	"log"

	"github.com/TeenBanner/Inventory_system/database"
	"github.com/TeenBanner/Inventory_system/migrations"
)

func main() {
	database.NewPostgresConnection()
	db := database.Pool()

	migrator := migrations.NewMigrator(db)

	if err := migrator.Migrate(db); err != nil {
		log.Fatalf("Error al realizar las migraciones. %v", err)
	}

}
