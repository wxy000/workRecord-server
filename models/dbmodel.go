package models

import (
	"server/globals"

	"gorm.io/gorm"
)

// 用户表
type Users struct {
	Userid   string `json:"userid" gorm:"size:20;not null;unique;primary_key;comment:工号"`
	Username string `json:"username" gorm:"size:20;not null;unique;comment:用户名-用于登录系统"`
	Realname string `json:"realname" gorm:"size:100;not null;comment:真实姓名"`
	Password string `json:"-" gorm:"size:100;not null;default:123456;comment:密码"`
	Phone    string `json:"phone" gorm:"size:100;comment:手机号码"`
	Email    string `json:"email" gorm:"size:100;comment:邮箱"`
	Gender   string `json:"gender" gorm:"size:1;not null;default:m;comment:性别-m男w女"`
	Mark     string `json:"mark" gorm:"size:1000;comment:备注"`
}

// 产品表
type Products struct {
	Productid   string `json:"productid" gorm:"size:10;not null;unique;primary_key;comment:产品编号"`
	Productname string `json:"productname" gorm:"size:100;not null;comment:产品名称"`
}

// 问题分类主表
type Issuemains struct {
	Mainid   string `json:"mainid" gorm:"size:20;not null;unique;primary_key"`
	Mainname string `json:"mainname" gorm:"size:100;not null"`
}

// 问题分类明细表
type Issuedetails struct {
	Detailid   string `json:"detailid" gorm:"size:20;not null;unique;primary_key"`
	Detailname string `json:"detailname" gorm:"size:100;not null"`
	Mainid     string `json:"mainid" gorm:"size:20;not null;comment:所属大类"`
}

// 客户表
type Customers struct {
	Customerid      string `json:"customerid" gorm:"size:20;not null;unique;primary_key;comment:客户id"`
	Customername    string `json:"customername" gorm:"size:100;not null;comment:客户名称"`
	Customeraddress string `json:"customeraddress" gorm:"size:100;comment:地址"`
	Cname           string `json:"cname" gorm:"size:100;comment:简称"`
}

// 记录表
type Record struct {
	gorm.Model
	Customerid         string  `json:"customerid" gorm:"size:20;not null;comment:客户编号"`
	Number             string  `json:"number" gorm:"size:100;not null;comment:作业编号"`
	Title              string  `json:"title" gorm:"size:100;not null;comment:问题标题"`
	Content            string  `json:"content" gorm:"size:2000;not null;comment:问题描述"`
	Feedbackid         string  `json:"feedbackid" gorm:"size:20;not null;comment:反馈人"`
	Feedbackdate       string  `json:"feedbackdate" gorm:"type:datetime;not null;comment:反馈时间"`
	Productid          string  `json:"productid" gorm:"size:10;not null;comment:产品"`
	Urgent             string  `json:"urgent" gorm:"size:1;not null;default:N;comment:是否紧急-Y紧急N不紧急"`
	Issuemainid        string  `json:"issuemainid" gorm:"size:20;not null;comment:问题分类"`
	Issuedetailid      string  `json:"issuedetailid" gorm:"size:20;not null;comment:问题类型"`
	Handlerid          string  `json:"handlerid" gorm:"size:20;not null;comment:处理人"`
	Handleestimatetime float32 `json:"handleestimatetime" gorm:"type:decimal(10,2);not null;default:0;comment:预计处理时长"`
	Handleactualtime   float32 `json:"handleactualtime" gorm:"type:decimal(10,2);not null;default:0;comment:实际处理时长"`
	Handlereply        string  `json:"handlereply" gorm:"size:2000;not null;comment:处理回复"`
	Casestatus         uint    `json:"casestatus" gorm:"not null;default:2;comment:案件状态-1.未结案，2.已结案"`
	Onsite             uint    `json:"onsite" gorm:"not null;default:2;comment:是否现场已处理-1.是，2.否"`
	Closetime          string  `json:"closetime" gorm:"type:datetime;not null;comment:结案时间"`
	Bugpeopleid        string  `json:"bugpeopleid" gorm:"size:20;comment:客制bug负责人"`
	Mark               string  `json:"mark" gorm:"size:2000;comment:备注"`
}

func Setup() {
	autoMigrate(&Users{}, &Products{}, &Issuemains{}, &Issuedetails{}, &Customers{}, &Record{})
}

// 自动迁移
func autoMigrate(tables ...interface{}) {
	// 创建表时添加后缀
	globals.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(tables...)
	// AutoMigrate 会创建表、缺失的外键、约束、列和索引。
	// 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。
}
