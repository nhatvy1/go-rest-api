package user_module

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, uc *UserController) {
	user := r.Group("/users")
	{
		user.GET("/", uc.GetList)
		user.GET("/:id", uc.GetUserById)
		user.POST("", uc.CreateUser)
	}
}
