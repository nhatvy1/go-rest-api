package initialize

import (
	user_module "user-service/internal/module/user"

	"gorm.io/gorm"
)

type Controllers struct {
	UserController *user_module.UserController
}

func InitAllModule(db *gorm.DB) *Controllers {
	return &Controllers{
		UserController: user_module.InitUserModule(db),
	}
}
