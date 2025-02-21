package services

import (
	"user-service/internal/models"
	"user-service/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(id int) error
}

type userService struct {
	userRepository repository.UserRepository
}

// CreateUser implements UserService.
func (u *userService) CreateUser(user *models.User) error {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetAllUsers implements UserService.
func (u *userService) GetAllUsers() ([]models.User, error) {
	panic("unimplemented")
}

// GetUserById implements UserService.
func (u *userService) GetUserById(id int) (*models.User, error) {
	panic("unimplemented")
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
