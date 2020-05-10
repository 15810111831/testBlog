package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"test.blog.com/testBlog/common"
	"test.blog.com/testBlog/model"
	"test.blog.com/testBlog/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		// 验证token是否需存在或是否以Bearer开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(ctx, http.StatusUnauthorized, 401, "权限不足", nil)
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		// 验证token是否解析错误，或者token过期
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, "权限不足", nil)
			ctx.Abort()
			return
		}

		// 获取token中的userid
		var user = model.User{}
		DB := common.DB
		DB.First(&user, claims.UserId)

		if user.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, "权限不足", nil)
			ctx.Abort()
			return
		}

		// 如果存在user则写入上下文中
		ctx.Set("user", user)
		ctx.Next()
	}
}
