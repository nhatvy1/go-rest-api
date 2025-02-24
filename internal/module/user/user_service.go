package user_module

import user_models "user-service/internal/module/user/models"

type UserService interface {
	GetAllUsers() string
	GetUserById(id int) (*user_models.User, error)
	CreateUser() string
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

func (u *userService) CreateUser() string {
	return "Chức năng tạo người dùng đang phát triển"
}

func (u *userService) DeleteUser() string {
	return "Chức năng xóa người dùng đang phát triển"
}

func (u *userService) GetAllUsers() string {
	return "Lấy danh sách người dùng đang phát triển"
}

func (u *userService) GetUserById(id int) (*user_models.User, error) {
	user, err := u.userRepository.FindById(1)
	if err != nil {
		return nil, err
	}
	return user, nil
}
