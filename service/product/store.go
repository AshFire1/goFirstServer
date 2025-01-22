package product

import (
	"database/sql"

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
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
		&product.Image,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
