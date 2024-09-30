package repository

import (
	"boilerplate/lib/database/entity"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id int) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserById(id int) (*entity.User, error) {
	user := entity.User{}
	result := r.db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}
