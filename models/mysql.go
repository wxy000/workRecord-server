package models

import (
	"log"
	"server/globals"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库初始化
func InitMysql() {
	log.Println("数据库初始化。。。")

	dsn := globals.Confok.Mysql.Dsn
	// log.Printf("dsn--------------------: %v\n", config.Confok.Database)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("打开mysql失败,", err)
		panic(err)
	}
	globals.DB = d
	sqlDB, _ := globals.DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(globals.Confok.Mysql.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(globals.Confok.Mysql.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
