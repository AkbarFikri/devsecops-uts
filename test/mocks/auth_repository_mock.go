// test/service/mocks/user_repository_mock.go
package mocks

import (
	"github.com/AkbarFikri/devsecops-uts/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUser() ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *UserRepositoryMock) GetUserByUsername(username string) (entity.User, error) {
	args := m.Called(username)
	return args.Get(0).(entity.User), args.Error(1)
}
