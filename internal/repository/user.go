package repository

import (
	"user-service/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id uint) (*models.User, error)
	Save(user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

// Save implements UserRepository.
func (u *userRepository) Save(user *models.User) error {
	panic("unimplemented")
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id uint) error {
	panic("unimplemented")
}

// FindAll implements UserRepository.
func (u *userRepository) FindAll() ([]models.User, error) {
	return []models.User{}, nil
}

// FindById implements UserRepository.
func (u *userRepository) FindById(id uint) (*models.User, error) {
	return &models.User{}, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
