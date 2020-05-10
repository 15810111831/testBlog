package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"test.blog.com/testBlog/common"
	"test.blog.com/testBlog/dto"
	"test.blog.com/testBlog/model"
	"test.blog.com/testBlog/response"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	fmt.Println(len(phone))
	if len(phone) != 11 {
		response.Fail(ctx, "手机号必须是11位", nil)
		return
	}

	if len(password) < 6 {
		response.Fail(ctx, "密码至少6位", nil)
		return
	}

	if isPhone(DB, phone) {
		response.Fail(ctx, "该手机号已被注册", nil)
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(ctx, "密码加密错误", nil)
		return
	}

	user := model.User{Name: name, Password: string(pwd[:]), Phone: phone}
	DB.Create(&user)
	response.Success(ctx, "创建成功", nil)
}

func isPhone(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Login(ctx *gin.Context) {
	DB := common.DB

	// 获取信息
	phone := ctx.PostForm("phone")
	passowrd := ctx.PostForm("password")

	// 获取user
	var user model.User
	DB.Where("phone = ?", phone).First(&user)

	if user.ID == 0 {
		response.Fail(ctx, "该用户为注册", nil)
		return
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passowrd)); err != nil {
		response.Fail(ctx, "密码不正确", nil)
		return
	}

	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Fail(ctx, "该用户为注册", nil)
		response.Response(ctx, http.StatusInternalServerError, 500, "获取token错误", nil)
		return
	}

	// 返回登录信息
	response.Success(ctx, "登陆成功", gin.H{
		"token": token,
	})

	return
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
	return
}
