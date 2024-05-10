package database

import (
	"database/sql"
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
	})
}

func Pool() *sql.DB {
	return db
}
