package repository

import (
	"database/sql"
	"errors"

	"github.com/royxu/simplegin/v2/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) GetUser(id int) (*model.User, error) {
	rows, err := ur.DB.Query(
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

func (ur *UserRepository) CreateUser(user *model.User) error {
	_, err := ur.DB.Exec(
		`INSERT INTO users (username, email) VALUES ($1, $2)`,
		user.Username, user.Email)

	if err != nil {
		return errors.New("user already exists")
	}

	return nil
}

func (ur *UserRepository) GetUsers() (*[]model.User, error) {
	rows, err := ur.DB.Query(
		`SELECT id, username, email FROM users`,
	)

	if err != nil {
		return nil, errors.New("Error fetching users")
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
		)
		users = append(users, user)
	}

	return &users, nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.DB.Exec(
		`DELETE FROM users WHERE id = $1`,
		id)

	if err != nil {
		return errors.New("Error deleting user")
	}

	return nil
}
