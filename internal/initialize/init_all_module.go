package initialize

import (
	"user-service/internal/controller"

	"gorm.io/gorm"
)

type Controllers struct {
	UserController *controller.UserController
}

func InitAllModule(db *gorm.DB) *Controllers {
	return &Controllers{
		UserController: InitUserModule(db),
	}
}
