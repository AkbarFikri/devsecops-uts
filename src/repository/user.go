package repository

import (
	"github.com/AkbarFikri/devsecops-uts/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() ([]entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) UserRepository {
	return &userRepository{
		db: gorm,
	}
}

func (a userRepository) GetUser() ([]entity.User, error) {
	var users []entity.User
	if err := a.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (a userRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
