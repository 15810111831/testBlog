package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, msg string, data gin.H) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

func Success(ctx *gin.Context, msg string, data gin.H) {
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": msg, "data": data})
}

func Fail(ctx *gin.Context, msg string, data gin.H) {
	ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": msg, "data": data})
}
