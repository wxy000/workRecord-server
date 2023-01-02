package controllers

import (
	"server/common"
	"server/models"

	"github.com/gin-gonic/gin"
)

// 一定日期范围内的时长变化曲线
func GetAnalysis1(c *gin.Context) {
	user, _ := c.Get("user")
	handlerid := user.(models.Users).Username
	feedbackdatestart := c.DefaultQuery("feedbackdatestart", "1900-01-01")
	if feedbackdatestart == "" {
		feedbackdatestart = "1900-01-01"
	}
	feedbackdatestart = feedbackdatestart + " 00:00:00"
	feedbackdateend := c.DefaultQuery("feedbackdateend", "3000-12-31")
	if feedbackdateend == "" {
		feedbackdateend = "3000-12-31"
	}
	feedbackdateend = feedbackdateend + " 23:59:59"
	succ, analysisRecordList1, count := models.GetAnalysisRecordList1(handlerid, feedbackdatestart, feedbackdateend)
	if succ {
		var date1s []string
		var es1s []float32
		var ac1s []float32
		for i := 0; i < len(*analysisRecordList1); i++ {
			date1s = append(date1s, (*analysisRecordList1)[i].Feedbackdate)
			es1s = append(es1s, (*analysisRecordList1)[i].Handleestimatetime)
			ac1s = append(ac1s, (*analysisRecordList1)[i].Handleactualtime)
		}
		common.OkWithDataC(count, gin.H{
			"date1": date1s,
			"y1":    es1s,
			"y2":    ac1s,
		}, c)
	} else {
		common.FailWithMsg("获取信息失败，请稍后重试", c)
	}
}
