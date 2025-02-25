package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponseData struct {
	Code   int         `json:"code"`
	Err    string      `json:"error"`
	Detail interface{} `json:"detail"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrResponse(c *gin.Context, code int, message string, detail interface{}) {
	c.JSON(http.StatusOK, ErrorResponseData{
		Code:   code,
		Err:    message,
		Detail: detail,
	})
}
