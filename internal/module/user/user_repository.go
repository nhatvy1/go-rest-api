package user_module

import (
	user_models "user-service/internal/module/user/models"
	"user-service/pkg/commons"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(p commons.PagingUser) ([]user_models.User, error)
	FindById(id uint) (*user_models.User, error)
	Save() string
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

func (u *userRepository) Save() string {
	return "save"
}

func (u *userRepository) Delete() string {
	return "delete"
}

func (u *userRepository) FindAll(p commons.PagingUser) ([]user_models.User, error) {
	var users []user_models.User
	offset := (p.Page - 1) * p.Limit

	if err := u.db.Order("id desc").
		Offset(offset).
		Limit(p.Limit).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) FindById(id uint) (*user_models.User, error) {
	var user user_models.User

	u.db.First(&user)
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
