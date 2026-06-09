package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/royxu/simplegin/v2/internal/model"
)

type UserRepository struct {
	DB *sql.DB
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

func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.DB.Exec(
		`DELETE FROM users WHERE id = $1`,
		id)

	if err != nil {
		return errors.New("Error deleting user")
	}

	return nil
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

func (ur *UserRepository) PatchUser(user *model.User) error {
	sets := []string{}
	args := []any{}
	i := 1
	if user.Username != "" {
		sets = append(sets, fmt.Sprintf("username = $%d", i))
		args = append(args, user.Username)
		i++
	}
	if user.Email != "" {
		sets = append(sets, fmt.Sprintf("email = $%d", i))
		args = append(args, user.Email)
		i++
	}
	query := fmt.Sprintf(
		"UPDATE users SET %s WHERE id = $%d",
		strings.Join(sets, ", "),
		i,
	)
	args = append(args, user.ID)

	res, err := ur.DB.Exec(query, args...)

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user not found")
	}
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (ur *UserRepository) PutUser(user *model.User) error {
	res, err := ur.DB.Exec(
		`UPDATE users SET username = $1, email = $2 WHERE id = $3`,
		user.Username, user.Email, user.ID)

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user not found")
	}
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
