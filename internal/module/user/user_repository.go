package user_module

import (
	"fmt"
	user_models "user-service/internal/module/user/models"
	"user-service/pkg/commons"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(p commons.PagingUser) ([]user_models.User, error)
	FindById(id int) (*user_models.User, error)
	FindByEmail(email string) (*user_models.User, error)
	Save(data *user_models.UserCreate) (*user_models.User, error)
	Delete() string
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (userRepo *userRepository) Save(data *user_models.UserCreate) (*user_models.User, error) {
	user := user_models.User{
		Email:     data.Email,
		Password:  data.Password,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	if err := userRepo.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *userRepository) Delete() string {
	return "delete"
}

func (userRepo *userRepository) FindAll(p commons.PagingUser) ([]user_models.User, error) {
	var users []user_models.User
	offset := (p.Page - 1) * p.Limit

	if err := userRepo.db.Order("id desc").
		Offset(offset).
		Limit(p.Limit).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepo *userRepository) FindById(id int) (*user_models.User, error) {
	var user user_models.User

	fmt.Println(id)
	if err := userRepo.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *userRepository) FindByEmail(email string) (*user_models.User, error) {
	user := user_models.User{}

	if err := userRepo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
