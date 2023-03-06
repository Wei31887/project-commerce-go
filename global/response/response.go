package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrResponse(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"error": err.Error(),
	})
}

func MsgResponse(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func SuccessResponse(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, obj)	
}