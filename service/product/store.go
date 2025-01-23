package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/AshFire1/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil

}

func scanRowsIntoProduct(row *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	fmt.Println(product)
	fmt.Print(err)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (s *Store) GetProductsByID(productIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)
	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)

	}
	return products, nil

}
func (s *Store) UpdateProduct(product types.Product) error {
	query := `UPDATE products SET name=?, description=?, price=?, quantity=? WHERE id=?`
	_, err := s.db.Exec(query, product.Name, product.Description, product.Price, product.Quantity, product.ID)
	return err
}
