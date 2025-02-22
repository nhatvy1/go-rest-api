package user_module

import (
	"gorm.io/gorm"
)

func InitUserModule(db *gorm.DB) *UserController {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userController := NewUserController(userService)

	return userController
}
