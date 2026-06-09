package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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

func TestRepositoryDeleteUser(t *testing.T) {
	userRepository := repo.UserRepository{
		DB: db,
	}

	newUser := model.User{
		Username: "testuser",
		Email:    "testuser@example.com",
	}
	err := userRepository.CreateUser(&newUser)
	require.NoError(t, err)

	users, err := userRepository.GetUsers()
	require.NoError(t, err)
	require.NotNil(t, users)
	assert.Equal(t, 1, len(*users), "User should be 1")

	err = userRepository.DeleteUser(int((*users)[0].ID))
	require.NoError(t, err)

	users, err = userRepository.GetUsers()
	require.NoError(t, err)
	require.NotNil(t, users)
	assert.Equal(t, 0, len(*users), "User should be 0")
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
		require.NoError(t, err)
	}

	users, err := userRepository.GetUsers()
	require.NoError(t, err)
	require.NotNil(t, users)
	assert.Equal(t, 2, len(*users), "User should be 2")

	for _, user := range *users {
		err := userRepository.DeleteUser(int(user.ID))
		require.NoError(t, err)
	}
}

func TestRepositoryPatchUser(t *testing.T) {
	userRepository := repo.UserRepository{
		DB: db,
	}

	newUser := model.User{
		Username: "testuser1",
		Email:    "testuser1@example.com",
	}
	err := userRepository.CreateUser(&newUser)
	require.NoError(t, err)

	users, err := userRepository.GetUsers()
	require.NoError(t, err)
	err = userRepository.PatchUser(&model.User{
		Username: "updateduser",
		ID:       int64((*users)[0].ID),
	})
	require.NoError(t, err)

	users, err = userRepository.GetUsers()
	require.NoError(t, err)
	require.NotNil(t, users)
	assert.Equal(t, "updateduser", (*users)[0].Username, "User should be updated")
	assert.Equal(t, "testuser1@example.com", (*users)[0].Email, "User should not be updated")

	err = userRepository.PatchUser(&model.User{
		Email: "updateduser@example.com",
		ID:    int64((*users)[0].ID),
	})
	require.NoError(t, err)

	users, err = userRepository.GetUsers()
	require.NoError(t, err)
	require.NotNil(t, users)
	assert.Equal(t, "updateduser", (*users)[0].Username, "User should not be updated")
	assert.Equal(t, "updateduser@example.com", (*users)[0].Email, "User should be updated")

	for _, user := range *users {
		err := userRepository.DeleteUser(int(user.ID))
		require.NoError(t, err)
	}
}

func TestRepositoryPutUser(t *testing.T) {
	userRepository := repo.UserRepository{
		DB: db,
	}

	newUser := model.User{
		Username: "testuser1",
		Email:    "testuser1@example.com",
	}
	err := userRepository.CreateUser(&newUser)
	require.NoError(t, err)

	users, err := userRepository.GetUsers()
	require.NoError(t, err)
	err = userRepository.PutUser(&model.User{
		Username: "updateduser",
		Email:    "updateduser@example.com",
		ID:       int64((*users)[0].ID),
	})
	require.NoError(t, err)

	users, err = userRepository.GetUsers()
	require.NoError(t, err)
	require.NotNil(t, users)
	assert.Equal(t, "updateduser", (*users)[0].Username, "User should be updated")
	assert.Equal(t, "updateduser@example.com", (*users)[0].Email, "User should be updated")

	for _, user := range *users {
		err := userRepository.DeleteUser(int(user.ID))
		require.NoError(t, err)
	}
}
