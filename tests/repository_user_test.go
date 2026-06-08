package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/royxu/simplegin/v2/configs"
	App "github.com/royxu/simplegin/v2/internal/app"
	model "github.com/royxu/simplegin/v2/internal/model"
	repo "github.com/royxu/simplegin/v2/internal/repository"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var config = configs.InitConfig("../configs/config.test.yaml")
	db = App.InitDB(&config)
	defer db.Close()

	code := m.Run()

	os.Exit(code)
}

func TestRepositoryGetUsers(t *testing.T) {
	userRepository := repo.UserRepository{
		DB: db,
	}

	newUsers := [2]model.User{
		{
			Username: "testuser1",
			Email:    "testuser1@example.com",
		},
		{
			Username: "testuser2",
			Email:    "testuser2@example.com",
		},
	}
	for _, newUser := range newUsers {
		err := userRepository.CreateUser(&newUser)
		if err != nil {
			t.Errorf("Create user failed: %v", err)
		}
	}

	users, err := userRepository.GetUsers()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(*users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(*users))
	}

	for _, user := range *users {
		err := userRepository.DeleteUser(int(user.ID))
		if err != nil {
			t.Errorf("Failed to delete user: %v", err)
		}
	}
}

func TestRepositoryDeleteUser(t *testing.T) {
	userRepository := repo.UserRepository{
		DB: db,
	}

	newUser := model.User{
		Username: "testuser",
		Email:    "testuser@example.com",
	}
	userRepository.CreateUser(&newUser)

	users, err := userRepository.GetUsers()
	if len(*users) != 1 {
		t.Errorf("Expected 1 users, got %d", len(*users))
	}

	err = userRepository.DeleteUser(int((*users)[0].ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	users, err = userRepository.GetUsers()
	if len(*users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(*users))
	}
}
