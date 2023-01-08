package controllers

import (
	"server/common"
	"server/models"
	"strconv"

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

// 一定日期范围内的客户时长分布
func GetAnalysis2(c *gin.Context) {
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
	succ, analysisRecordList2, count := models.GetAnalysisRecordList2(handlerid, feedbackdatestart, feedbackdateend)
	succ1, analysisRecordListSum, _ := models.GetAnalysisRecordListSum(handlerid, feedbackdatestart, feedbackdateend)
	if succ && succ1 {
		var es2s []map[string]string
		var ac2s []map[string]string
		var tbs []map[string]interface{}
		for i := 0; i < len(*analysisRecordList2); i++ {
			es2 := map[string]string{"name": (*analysisRecordList2)[i].Cname, "value": strconv.FormatFloat(float64((*analysisRecordList2)[i].Handleestimatetime), 'f', 2, 32)}
			ac2 := map[string]string{"name": (*analysisRecordList2)[i].Cname, "value": strconv.FormatFloat(float64((*analysisRecordList2)[i].Handleactualtime), 'f', 2, 32)}
			es2s = append(es2s, es2)
			ac2s = append(ac2s, ac2)
			// ---
			es := (*analysisRecordList2)[i].Handleestimatetime
			esv := (*analysisRecordList2)[i].Handleestimatetime / (*analysisRecordListSum)[0].Handleestimatetime
			ac := (*analysisRecordList2)[i].Handleactualtime
			acv := (*analysisRecordList2)[i].Handleactualtime / (*analysisRecordListSum)[0].Handleactualtime
			tb := map[string]interface{}{"name": (*analysisRecordList2)[i].Cname, "es": es, "esv": esv, "ac": ac, "acv": acv}
			tbs = append(tbs, tb)
		}
		common.OkWithDataC(count, gin.H{
			"es2s": es2s,
			"ac2s": ac2s,
			"tbs":  tbs,
		}, c)
	} else {
		common.FailWithMsg("获取信息失败，请稍后重试", c)
	}
}

// 一定日期范围内的问题类型时长分布
func GetAnalysis3(c *gin.Context) {
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
	succ, analysisRecordList3, count := models.GetAnalysisRecordList3(handlerid, feedbackdatestart, feedbackdateend)
	succ1, analysisRecordList3_1, _ := models.GetAnalysisRecordList3_1(handlerid, feedbackdatestart, feedbackdateend)
	succ2, analysisRecordListSum, _ := models.GetAnalysisRecordListSum(handlerid, feedbackdatestart, feedbackdateend)
	if succ && succ1 && succ2 {
		// 标准分类
		var es3sall []map[string]interface{}
		var es3sall_item_sz []map[string]interface{}
		for i := 0; i < len(*analysisRecordList3); i++ {
			es3sall_item_sz_item := map[string]interface{}{"name": (*analysisRecordList3)[i].Detailname, "value": (*analysisRecordList3)[i].Handleactualtime}
			if i == 0 {
				es3sall_item_sz = append(es3sall_item_sz, es3sall_item_sz_item)
			} else {
				if (*analysisRecordList3)[i].Mainname == (*analysisRecordList3)[i-1].Mainname {
					es3sall_item_sz = append(es3sall_item_sz, es3sall_item_sz_item)
				} else {
					es3sall_item := map[string]interface{}{"name": (*analysisRecordList3)[i-1].Mainname, "children": es3sall_item_sz}
					es3sall = append(es3sall, es3sall_item)

					es3sall_item_sz = []map[string]interface{}{}
					es3sall_item_sz = append(es3sall_item_sz, es3sall_item_sz_item)
				}
			}
			if i == len(*analysisRecordList3)-1 {
				es3sall_item := map[string]interface{}{"name": (*analysisRecordList3)[i].Mainname, "children": es3sall_item_sz}
				es3sall = append(es3sall, es3sall_item)
			}
		}
		// 自定义分类
		var es3s []map[string]interface{}
		var ac3s []map[string]interface{}
		var tbs []map[string]interface{}
		for i := 0; i < len(*analysisRecordList3_1); i++ {
			es3 := map[string]interface{}{"name": (*analysisRecordList3_1)[i].Classname, "value": (*analysisRecordList3_1)[i].Handleestimatetime}
			ac3 := map[string]interface{}{"name": (*analysisRecordList3_1)[i].Classname, "value": (*analysisRecordList3_1)[i].Handleactualtime}
			es3s = append(es3s, es3)
			ac3s = append(ac3s, ac3)
			// ---
			es := (*analysisRecordList3_1)[i].Handleestimatetime
			esv := (*analysisRecordList3_1)[i].Handleestimatetime / (*analysisRecordListSum)[0].Handleestimatetime
			ac := (*analysisRecordList3_1)[i].Handleactualtime
			acv := (*analysisRecordList3_1)[i].Handleactualtime / (*analysisRecordListSum)[0].Handleactualtime
			tb := map[string]interface{}{"name": (*analysisRecordList3_1)[i].Classname, "es": es, "esv": esv, "ac": ac, "acv": acv}
			tbs = append(tbs, tb)
		}
		common.OkWithDataC(count, gin.H{
			"es3sall": es3sall,
			"es3s":    es3s,
			"ac3s":    ac3s,
			"tbs":     tbs,
		}, c)
	} else {
		common.FailWithMsg("获取信息失败，请稍后重试", c)
	}
}
