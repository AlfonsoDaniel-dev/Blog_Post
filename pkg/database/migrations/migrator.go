package migrations

import (
	"database/sql"
	"fmt"
	"log"
)

type Migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) Migrator {
	return Migrator{
		db: db,
	}
}

func (m *Migrator) MigrateExtension(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("extencion de uuid migrada con exito")

	return nil
}

func (M *Migrator) MigrateRollTable(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("Migracion de roles completada")

	return nil
}

func (M *Migrator) MigrateTablePosts(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de la tabla posts completada")

	return nil
}

func (M *Migrator) MigrateTableUsers(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("Migracion de users completada")

	return nil
}

func (M *Migrator) migrateAdminTable(tx *sql.Tx, query string) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("Migracion de administradores completada")

	return nil
}

func (m *Migrator) Migrate(db *sql.DB) error {
	// migrar tabla usuarios
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := m.MigrateExtension(tx, SqlCreateUuidExtension); err != nil {
		tx.Rollback()
		return err
	}

	if err := m.MigrateTableUsers(tx, SqlCreateUserTable); err != nil {
		tx.Rollback()
		return err
	}

	if err := m.MigrateTablePosts(tx, SqlCreatePostTable); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("Migraciones realizadas")

	return tx.Commit()
}
