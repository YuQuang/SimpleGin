package repository

import (
	"database/sql"
	"errors"

	"github.com/royxu/simplegin/v2/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetUser(
	id int,
) (*model.User, error) {
	rows, err := r.DB.Query(
		`SELECT username, email FROM users WHERE id = $1`,
		id)

	if err != nil {
		return &model.User{}, errors.New("Query error")
	}
	if !rows.Next() {
		return &model.User{}, errors.New("user not found")
	}

	var user model.User
	rows.Scan(
		&user.Username,
		&user.Email,
	)

	return &user, nil
}

func (r *UserRepository) CreateUser(
	user *model.User,
) error {
	_, err := r.DB.Query(
		`INSERT INTO users (username, email) VALUES ($1, $2)`,
		user.Username, user.Email)

	if err != nil {
		return errors.New("user already exists")
	}

	return nil
}
