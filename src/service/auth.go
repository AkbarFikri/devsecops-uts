package service

import (
	"errors"
	"github.com/AkbarFikri/devsecops-uts/src/repository"
)

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a authService) Login(username, password string) (string, error) {
	user, err := a.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("wrong password")
	}

	return user.Username, nil
}
