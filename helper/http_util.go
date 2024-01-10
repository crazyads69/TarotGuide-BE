package helper

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPError(ctx *gin.Context, status int, err error, message string) {
	er := gin.H{
		"status":  status,
		"error":   err.Error(),
		"message": message,
	}
	ctx.AbortWithStatusJSON(status, er)
}

func NewHTTPResponse(ctx *gin.Context, status int, data interface{}, message string) {
	rsp := gin.H{
		"status":  status,
		"data":    data,
		"message": message,
	}
	ctx.JSON(status, rsp)
}
