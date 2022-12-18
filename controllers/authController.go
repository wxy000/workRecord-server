package controllers

import (
	"server/common"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 登录
func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 校验用户名和密码是否正确
	user, err := models.LoginByUsername(username, password)
	if err != nil {
		common.FailWithMsg(err.Error(), c)
	} else {
		// 生成Token
		tokenString, _ := utils.GenToken(*user)
		common.OkWithData(gin.H{"token": tokenString}, c)
	}
}

// 根据token获取用户信息
func GetUserInfo(c *gin.Context) {
	user, err := c.Get("user")
	if !err {
		common.FailWithMsg("获取用户信息失败", c)
	} else {
		common.OkWithData(user, c)
	}
}

// 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	user := models.Users{
		Userid:   c.PostForm("userid"),
		Username: c.PostForm("username"),
		Realname: c.PostForm("realname"),
		Phone:    c.PostForm("phone"),
		Email:    c.PostForm("email"),
		Gender:   c.PostForm("gender"),
		Mark:     c.PostForm("mark"),
	}
	if err := models.UpdateUserInfoByID(user); err != nil {
		common.FailWithMsg(err.Error(), c)
	} else {
		common.OkWithMsg("用户信息更新成功", c)
	}
}

// 修改密码
func UpdatePassword(c *gin.Context) {
	user, _ := c.Get("user")
	userid := user.(models.Users).Userid
	oldpw := c.PostForm("oldpassword")
	password := c.PostForm("password")
	if err := models.UpdatePasswordByID(userid, oldpw, password); err != nil {
		common.FailWithMsg(err.Error(), c)
	} else {
		common.OkWithMsg("密码更新成功", c)
	}
}
