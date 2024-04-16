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
func (m *Migrator) migrateproduct(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de producto creada")

	return nil
}
func (m *Migrator) migrateorder(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de order realizada")

	return nil
}

func (m *Migrator) migrateUser(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de User realizada")

	return nil
}

func (m *Migrator) migrateUserAdmin(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de administradores completada")

	return nil
}

func (m *Migrator) Migrate(db *sql.DB) error {
	// migrar tabla usuarios
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := m.migrateUser(tx, psqlUser); err != nil {
		tx.Rollback()
		return err
	}

	if err := m.migrateproduct(tx, psqlProduct); err != nil {
		tx.Rollback()
		return err
	}

	if err := m.migrateorder(tx, psqlOrder); err != nil {
		tx.Rollback()
		return err
	}

	if err := m.migrateUserAdmin(tx, psqladmin); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("Migraciones realizadas")

	return tx.Commit()
}
