package migrations

import (
	"database/sql"
	"fmt"
)

const psqlMigrateproduct = `CREATE TABLE IF NOT EXISTS products(
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	name varchar(60) NOT NULL,
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	Upadated_at TIMESTAMP,
	CONSTRAINT products_id_pk PRIMARY KEY (id)
)`

type psqlproduct struct {
	db *sql.DB
}

func NewPsqlproduct(db *sql.DB) psqlproduct {
	return psqlproduct{
		db: db,
	}
}

func (p *psqlproduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateproduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	fmt.Println("Migracion de producto creada")

	return nil
}
