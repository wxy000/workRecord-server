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
		analysis := api.Group("/analysis")
		{
			// 一定日期范围内的时长变化曲线
			analysis.GET("/my/getAnalysis1", middleware.JWTAuthMiddleware(), controllers.GetAnalysis1)
			// 一定日期范围内的客户时长分布
			analysis.GET("/my/getAnalysis2", middleware.JWTAuthMiddleware(), controllers.GetAnalysis2)
			// 一定日期范围内的问题类型时长分布
			analysis.GET("/my/getAnalysis3", middleware.JWTAuthMiddleware(), controllers.GetAnalysis3)
			// 一定日期范围内的数据明细
			analysis.GET("/my/getAnalysis4", middleware.JWTAuthMiddleware(), controllers.GetAnalysis4)
			// 一定日期范围内的-签单/返工/问题处理
			analysis.GET("/my/getAnalysis5", middleware.JWTAuthMiddleware(), controllers.GetAnalysis5)
			// 一定日期范围内的-按月按客户统计
			analysis.GET("/my/getAnalysis6", middleware.JWTAuthMiddleware(), controllers.GetAnalysis6)
		}
	}

}
