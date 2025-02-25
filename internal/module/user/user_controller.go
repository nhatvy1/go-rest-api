package user_module

import (
	"fmt"
	"net/http"
	"strconv"
	user_models "user-service/internal/module/user/models"
	"user-service/pkg/response"
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
		response.ErrResponse(c, response.ErrCodeParamInvalid, "", nil)
	}

	user, err := uc.us.GetUserById(id)
	if err != nil {
		response.ErrResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	response.SuccessResponse(c, response.ErrCodeSuccess, user)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	userCreate := user_models.UserCreate{}

	if err := c.ShouldBindJSON(&userCreate); err != nil {
		fmt.Printf("%+v\n", userCreate)
		errMsgs := validations.ValidateRequest(err, userCreate)
		if errMsgs != nil {
			response.ErrResponse(c, http.StatusBadRequest, "", errMsgs)
			return
		}
		return
	}

	user, err := uc.us.CreateUser(&userCreate)
	if err != nil {
		response.ErrResponse(c, http.StatusConflict, "Email đã tồn tại", "")
		return
	}

	response.SuccessResponse(c, response.ErrCodeSuccess, user)
}

func (uc *UserController) GetList(c *gin.Context) {
	paging := user_models.Paging{}

	if err := c.ShouldBind(&paging); err != nil {
		response.ErrResponse(c, http.StatusBadRequest, "", err.Error())
		return
	}
	fmt.Printf("%+v\n", paging)

	response.SuccessResponse(c, response.ErrCodeSuccess, "list users")
}
