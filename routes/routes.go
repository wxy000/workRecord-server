package routes

import (
	"log"
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

// 这里的匿名函数后续要写到controllers里面
func RegisterRoute(c *gin.Engine) {

	log.Println("路由信息初始化。。。")

	// 使用跨域
	c.Use(middleware.Cors())

	api := c.Group("/api")
	{
		// 登陆注册；获取token
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.AuthHandler)
			auth.GET("/getUserInfo", middleware.JWTAuthMiddleware(), controllers.GetUserInfo)
			auth.POST("/updateUserInfo", middleware.JWTAuthMiddleware(), controllers.UpdateUserInfo)
			auth.POST("/updatePassword", middleware.JWTAuthMiddleware(), controllers.UpdatePassword)
		}
		workRecord := api.Group("/workRecord")
		{
			workRecord.GET("/wr/getRecordByHandlerid", middleware.JWTAuthMiddleware(), controllers.GetRecordByHandlerid)
			workRecord.GET("/wr/downloadRecordTemplate", controllers.DownloadRecordTemplate)
			workRecord.POST("/wr/importData", middleware.JWTAuthMiddleware(), controllers.ImportData)
		}
	}

}
