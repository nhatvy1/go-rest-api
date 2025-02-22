package user_module

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController *UserController) {
	user := r.Group("/users")
	{
		user.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user",
			})
		})
		user.GET("/:id", userController.GetUserById)
	}
}
