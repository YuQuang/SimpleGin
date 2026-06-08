package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/royxu/simplegin/v2/configs"
	App "github.com/royxu/simplegin/v2/internal/app"
	"github.com/royxu/simplegin/v2/internal/repository"
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
	userRepository := repository.UserRepository{
		DB: db,
	}
	users, _ := userRepository.GetUsers()

	if users == nil {
		t.Fatalf("Expected users, got nil")
	}
	if len(*users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(*users))
	}
}
