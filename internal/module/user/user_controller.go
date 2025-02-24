package user_module

import (
	"fmt"
	"net/http"
	"strconv"
	user_models "user-service/internal/module/user/models"
	"user-service/pkg/validations"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us UserService
}

func NewUserController(us UserService) *UserController {
	return &UserController{
		us: us,
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
	}
	user, err := uc.us.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	userCreate := user_models.UserCreate{}

	if err := c.ShouldBindJSON(&userCreate); err != nil {
		errMsgs := validations.ValidateRequest(err, userCreate)
		if errMsgs != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errMsgs})
			return
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userCreate,
	})
}

func (uc *UserController) GetList(c *gin.Context) {
	paging := user_models.Paging{}

	if err := c.ShouldBind(&paging); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	fmt.Printf("%+v\n", paging)

	c.JSON(http.StatusOK, gin.H{
		"data": "list user",
	})
}
