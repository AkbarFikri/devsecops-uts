// test/service/auth_service_test.go
package service_test

import (
	"errors"
	"github.com/AkbarFikri/devsecops-uts/entity"
	"testing"

	"github.com/AkbarFikri/devsecops-uts/src/service"
	"github.com/AkbarFikri/devsecops-uts/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		mockRepo := new(mocks.UserRepositoryMock)
		authService := service.NewAuthService(mockRepo)

		testUser := entity.User{
			Username: "admin",
			Password: "password123",
		}

		mockRepo.On("GetUserByUsername", "admin").Return(testUser, nil)

		username, err := authService.Login("admin", "password123")

		assert.NoError(t, err)
		assert.Equal(t, "admin", username)
		mockRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockRepo := new(mocks.UserRepositoryMock)
		authService := service.NewAuthService(mockRepo)

		mockRepo.On("GetUserByUsername", "unknown").Return(entity.User{}, errors.New("user not found"))

		username, err := authService.Login("unknown", "password")

		assert.Error(t, err)
		assert.Equal(t, "", username)
		assert.EqualError(t, err, "user not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Wrong Password", func(t *testing.T) {
		mockRepo := new(mocks.UserRepositoryMock)
		authService := service.NewAuthService(mockRepo)

		testUser := entity.User{
			Username: "admin",
			Password: "password123",
		}

		mockRepo.On("GetUserByUsername", "admin").Return(testUser, nil)

		username, err := authService.Login("admin", "wrongpass")

		assert.Error(t, err)
		assert.Equal(t, "", username)
		assert.EqualError(t, err, "wrong password")
		mockRepo.AssertExpectations(t)
	})
}
