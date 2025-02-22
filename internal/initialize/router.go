package initialize

import (
	user_module "user-service/internal/module/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	controllers := InitAllModule(db)

	v1 := r.Group("/api/v1")
	{
		user_module.UserRoutes(v1, controllers.UserController)
	}

	return r
}
