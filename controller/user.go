package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"test.blog.com/testBlog/common"
	"test.blog.com/testBlog/model"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	fmt.Println(len(phone))
	if len(phone) != 11 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "手机号必须是11位",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码至少6位",
		})
		return
	}

	if isPhone(DB, phone) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "手机号已被注册",
		})
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码加密错误",
		})
		return
	}

	user := model.User{Name: name, Password: string(pwd[:]), Phone: phone}
	DB.Create(&user)
	ctx.JSON(201, "创建成功")
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "该用户未注册",
		})
		return
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passowrd)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码不正确",
		})
		return
	}

	token, err := common.ReleaseToken(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取token错误",
		})
		return
	}

	// 返回登录信息
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
	return
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": user},
	})
	return
}
