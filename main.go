package main

import (
	"server/conf"
	"server/globals"
	"server/models"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.InitConfigByViper()
	models.InitMysql()
	// 数据库初始化
	models.Setup()
	// 路由
	r := gin.Default()
	routes.RegisterRoute(r)
	r.Run(":" + globals.Confok.Api.Port)
}
