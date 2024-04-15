package migrations

import (
	"database/sql"
	"fmt"
)

type Migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) Migrator {
	return Migrator{
		db: db,
	}
}
func Migrateproduct(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	fmt.Println("Migracion de producto creada")

	return nil
}
func Migrateorder(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	stmt.Close()

	fmt.Println("Migracion de order realizada")

	return nil
}

func MigrateUser(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	fmt.Println("Migracion de User realizada")

	return nil
}

func (M *Migrator) Migrate(db *sql.DB) error {
	// migrar tabla usuarios
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := MigrateUser(tx, psqlUser); err != nil {
		tx.Rollback()
		return err
	}

	if err := Migrateproduct(tx, psqlProduct); err != nil {
		tx.Rollback()
		return err
	}

	if err := Migrateorder(tx, psqlOrder); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("Transaccion de migracion realizada")

	return tx.Commit()
}
