package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresConnection() {
	once.Do(func() {
		var err error

		db, err = sql.Open("postgres", "")
		if err != nil {
			log.Fatalf("no se pudo conectar a la DB", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo hacer ping a la Base de datos. ERR: %v", err)
		}

		fmt.Println("conectado a la base de datos")
	})
}

func Pool() *sql.DB {
	return db
}
