package initialize

import (
	"user-service/internal/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	controllers := InitAllModule(db)

	v1 := r.Group("/api/v1")
	{
		routes.UserRoutes(v1, controllers.UserController)
	}

	return r
}
