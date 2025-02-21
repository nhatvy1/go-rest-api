package initialize

import (
	"user-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		routes.UserRoutes(v1)
		routes.RoleRoutes(v1)
	}

	return r
}
