package user

import (
	"database/sql"
	"fmt"

	"github.com/AshFire1/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// implement the logic to fetch the user from the database
	rows, err := s.db.Query("SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *Store) GetUserByID(id int) (*types.User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE id=$1", id)

	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *Store) CreateUser(user types.User) error {
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)"
	_, err := s.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
