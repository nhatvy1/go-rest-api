package routes

import (
	"net/http"
	"user-service/internal/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController *controller.UserController) {
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
