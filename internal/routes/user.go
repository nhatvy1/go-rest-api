package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("/users")
	{
		user.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user",
			})
		})
		user.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user post",
			})
		})
	}
}
