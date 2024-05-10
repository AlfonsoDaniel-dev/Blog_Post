package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

func CreateConnection(conStr string) {
	once.Do(func() {
		db, err = sql.Open("postgres", conStr)
		if err != nil {
			log.Fatalf("Cannot connect to DB. ERR: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("Cannot ping DB. ERR %v", err)
		}
	})
}

func Pool() *sql.DB {
	return db
}
