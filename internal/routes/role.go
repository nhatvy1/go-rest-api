package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(r *gin.RouterGroup) {
	role := r.Group("/roles")
	{
		role.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "role",
			})
		})
		role.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "role post",
			})
		})
	}
}
