package migrations

import (
	"database/sql"
	"fmt"
)

const psqlMigrateUser = `CREATE TABLE IF NOT EXISTS users(
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	name varchar(100) NOT NULL,
	email varchar(50) NOT NULL,
	password varchar(150) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	Updated_at TIMESTAMP,
	CONSTRAINT users_id_pk PRIMARY KEY (id),
	CONSTRAINT users_email_uq UNIQUE (email)
)`

type psqlUser struct {
	db *sql.DB
}

func (u *psqlUser) Migrate() error {
	stmt, err := u.db.Prepare(psqlMigrateUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	fmt.Println("Migracion de la tabla usuarios completada")

	return nil
}
