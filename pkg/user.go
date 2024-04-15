package pkg

import "database/sql"

type User struct {
	db *sql.DB
}
