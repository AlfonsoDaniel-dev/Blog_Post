package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/TeenBanner/Inventory_system/models"
	"github.com/google/uuid"
)

const sqlCreateProduct = `INSERT INTO products(owner_id, name, price) VALUES ($1, $2, 3$) RETURNING id`
const sqlGetProduct = `SELECT (id, owner_id, name, price, created_at) FROM products WHERE name = $2`
const sqlGetAllProducts = `SELECT (id, owner_id, name, price, created_at) FROM products`
const sqlDeleteProduct = `DELETE * FROM PRODUCTS WHERE id = $1`
const sqlUpdateProductName = `UPDATE products SET name = $1 WHERE id = $2`
const sqlUpdateProductPrice = `UPDATE products SET price = $1 WHERE id = $2 `

type psqlprduct struct {
	DB *sql.DB
}

func NewPsqlProduct(db *sql.DB) psqlprduct {
	return psqlprduct{
		DB: db,
	}
}

func (p *psqlprduct) CreateProduct(owner int, product models.Product) error {
	stmt, err := p.DB.Prepare(sqlCreateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	product.ID = uuid.New()

	row, err := stmt.Exec(product.OwnerID, product.Name, product.Price)
	if err != nil {
		return err
	}

	LastInsertID, err := row.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Printf("Producto Creado. ID del ultimo registro insertado: %v", LastInsertID)

	return nil
}

func (p *psqlprduct) GetProduct(name string) (models.Product, error) {
	stmt, err := p.DB.Prepare(sqlGetProduct)
	if err != nil {
		return models.Product{}, err
	}

	defer stmt.Close()

	product := models.Product{}

	row, err := stmt.Query(name)
	if err != nil {
		return models.Product{}, errors.New("el producto que buscas no existe")
	}

	updatedAtNull := sql.NullTime{}

	err = row.Scan(&product.ID, &product.OwnerID, &product.Name, &product.Price, &product.Created_At, &updatedAtNull)
	product.Updated_at = updatedAtNull.Time

	if err != nil {
		return models.Product{}, err
	}

	fmt.Printf("producto obtenido por su nombre")

	return product, nil
}

func (p *psqlprduct) GetAllProducts() ([]models.Product, error) {
	stmt, err := p.DB.Prepare(sqlGetAllProducts)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.Product{}

		updatedAtNull := sql.NullTime{}

		err := rows.Scan(&product.ID, &product.OwnerID, &product.Name, &product.Price, &product.Created_At, &updatedAtNull)
		product.Updated_at = updatedAtNull.Time
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("se obtuvieron todos los productos existentes")

	return products, nil
}

func (p *psqlprduct) DeleteProduct(id int) error {
	stmt, err := p.DB.Prepare(sqlDeleteProduct)
	if err != nil {
		return err
	}

	row, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	rowsAff, err := row.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Producto Eliminado, Filas Afectadas: %v", rowsAff)

	return nil
}

func (p *psqlprduct) UpdateProductName(id int, product models.Product) error {
	stmt, err := p.DB.Prepare(sqlUpdateProductName)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row, err := stmt.Exec(product.Name)
	if err != nil {
		return err
	}

	rowsaff, err := row.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Producto actualizado. Filas afectadas: %v", rowsaff)

	return nil
}

func (p *psqlUser) UpdateProductPrice(id int, product models.Product) error {
	stmt, err := p.DB.Prepare(sqlUpdateProductPrice)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row, err := stmt.Exec(product.Price, id)
	if err != nil {
		return err
	}

	rowsAff, err := row.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Producto actualizado. Filas afectadas: %v", rowsAff)

	return err
}
