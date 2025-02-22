package initialize

import (
	"user-service/internal/controller"
	"user-service/internal/repository"
	"user-service/internal/services"

	"gorm.io/gorm"
)

func InitUserModule(db *gorm.DB) *controller.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	return userController
}
