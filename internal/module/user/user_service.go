package user_module

import (
	"fmt"
	user_models "user-service/internal/module/user/models"
	"user-service/pkg/response"
)

type UserService interface {
	GetAllUsers() string
	GetUserById(id int) (*user_models.User, error)
	CreateUser(uc *user_models.UserCreate) (int, error)
	DeleteUser() string
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) CreateUser(userCreate *user_models.UserCreate) (int, error) {
	checkUserEmail, _ := u.userRepository.FindByEmail(userCreate.Email)
	fmt.Printf("check: %+v\n", checkUserEmail)
	if checkUserEmail != nil {
		return response.ErrCodeNotFound, fmt.Errorf("%s", "Email đã tồn tại")
	}

	// u.userRepository.Save(userCreate)
	return 1, nil
}

func (u *userService) DeleteUser() string {
	return "Chức năng xóa người dùng đang phát triển"
}

func (u *userService) GetAllUsers() string {
	return "Lấy danh sách người dùng đang phát triển"
}

func (u *userService) GetUserById(id int) (*user_models.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
