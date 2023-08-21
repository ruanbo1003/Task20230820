package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response gin restful api response
type Response struct {
	ErrorCode int32       `json:"error_code"` // error code, 0 - success, non zero - fail
	Message   string      `json:"message"`    // error message, "" - success, non empty - the fail message
	Data      interface{} `json:"data"`       // response data
}

// Success successful response
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		ApiSuccess,
		"ok",
		data,
	})
}

// Fail non-successful response
func Fail(c *gin.Context, errorCode int32, errMessage string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		errMessage,
		nil,
	})
}
