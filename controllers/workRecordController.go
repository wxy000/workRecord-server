package controllers

import (
	"server/common"
	"server/models"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRecordByHandlerid(c *gin.Context) {
	user, _ := c.Get("user")
	handlerid := user.(models.Users).Userid
	succ, recordList, count := models.GetRecordByHandlerid(handlerid, c)
	if succ {
		for i := 0; i < len(*recordList); i++ {
			// 反馈时间
			t1, _ := time.Parse(time.RFC3339, (*recordList)[i].Feedbackdate)
			(*recordList)[i].Feedbackdate = t1.Format("2006-01-02 15:04:05")
			// 结案时间
			t2, _ := time.Parse(time.RFC3339, (*recordList)[i].Closetime)
			(*recordList)[i].Closetime = t2.Format("2006-01-02 15:04:05")
		}
		common.OkWithDataC(count, recordList, c)
	} else {
		common.FailWithMsg("获取信息失败，请稍后重试", c)
	}
}

func DownloadRecordTemplate(c *gin.Context) {
	filePath := "./data/导入excel模板.xlsx"
	utils.DownloadFile(c, filePath)
}

func ImportData(c *gin.Context) {
	succ, err := utils.ImportDataWR(c)
	if succ {
		common.OkWithMsg(err.Error(), c)
	} else {
		common.FailWithMsg(err.Error(), c)
	}
}
