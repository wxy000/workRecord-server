package controllers

import (
	"server/common"
	"server/models"
	"sort"
	"strconv"
	"time"

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

// 一定日期范围内的数据明细
func GetAnalysis4(c *gin.Context) {
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
	succ, analysisRecordList4, count := models.GetAnalysisRecordList4(handlerid, feedbackdatestart, feedbackdateend)
	if succ {
		for i := 0; i < len(*analysisRecordList4); i++ {
			// 反馈时间
			t1, _ := time.Parse(time.RFC3339, (*analysisRecordList4)[i].Feedbackdate)
			(*analysisRecordList4)[i].Feedbackdate = t1.Format("2006-01-02 15:04:05")
			// 结案时间
			t2, _ := time.Parse(time.RFC3339, (*analysisRecordList4)[i].Closetime)
			(*analysisRecordList4)[i].Closetime = t2.Format("2006-01-02 15:04:05")
		}
		common.OkWithDataC(count, analysisRecordList4, c)
	} else {
		common.FailWithMsg("获取信息失败，请稍后重试", c)
	}
}

// 一定日期范围内的-签单/返工/问题处理
func GetAnalysis5(c *gin.Context) {
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
	succ, analysisRecordList5, count := models.GetAnalysisRecordList5(handlerid, feedbackdatestart, feedbackdateend)
	if succ {
		// 预计
		var es5s2s [][]map[string]interface{}
		var es5s2 []map[string]interface{}
		// 实际
		var ac5s2s [][]map[string]interface{}
		var ac5s2 []map[string]interface{}
		for i := 0; i < len(*analysisRecordList5); i++ {
			if i == 0 {
				es5s2tmp := map[string]interface{}{(*analysisRecordList5)[i].Ny: (*analysisRecordList5)[i].Handleestimatetime}
				es5s2 = append(es5s2, es5s2tmp)
				ac5s2tmp := map[string]interface{}{(*analysisRecordList5)[i].Ny: (*analysisRecordList5)[i].Handleactualtime}
				ac5s2 = append(ac5s2, ac5s2tmp)
			} else {
				if (*analysisRecordList5)[i].Typea == (*analysisRecordList5)[i-1].Typea {
					es5s2tmp := map[string]interface{}{(*analysisRecordList5)[i].Ny: (*analysisRecordList5)[i].Handleestimatetime}
					es5s2 = append(es5s2, es5s2tmp)
					ac5s2tmp := map[string]interface{}{(*analysisRecordList5)[i].Ny: (*analysisRecordList5)[i].Handleactualtime}
					ac5s2 = append(ac5s2, ac5s2tmp)
				} else {
					if (*analysisRecordList5)[i-1].Typea == "1" {
						es5s2tmp := map[string]interface{}{"name": "客制签单时数(预计)"}
						es5s2 = append(es5s2, es5s2tmp)
						ac5s2tmp := map[string]interface{}{"name": "客制签单时数(实际)"}
						ac5s2 = append(ac5s2, ac5s2tmp)
					} else if (*analysisRecordList5)[i-1].Typea == "2" {
						es5s2tmp := map[string]interface{}{"name": "返工赠送时数(预计)"}
						es5s2 = append(es5s2, es5s2tmp)
						ac5s2tmp := map[string]interface{}{"name": "返工赠送时数(实际)"}
						ac5s2 = append(ac5s2, ac5s2tmp)
					} else if (*analysisRecordList5)[i-1].Typea == "3" {
						es5s2tmp := map[string]interface{}{"name": "问题处理时数(预计)"}
						es5s2 = append(es5s2, es5s2tmp)
						ac5s2tmp := map[string]interface{}{"name": "问题处理时数(实际)"}
						ac5s2 = append(ac5s2, ac5s2tmp)
					}
					es5s2s = append(es5s2s, es5s2)
					es5s2 = []map[string]interface{}{}
					es5s2tmp := map[string]interface{}{(*analysisRecordList5)[i].Ny: (*analysisRecordList5)[i].Handleestimatetime}
					es5s2 = append(es5s2, es5s2tmp)
					ac5s2s = append(ac5s2s, ac5s2)
					ac5s2 = []map[string]interface{}{}
					ac5s2tmp := map[string]interface{}{(*analysisRecordList5)[i].Ny: (*analysisRecordList5)[i].Handleactualtime}
					ac5s2 = append(ac5s2, ac5s2tmp)
				}
			}
			if i == len(*analysisRecordList5)-1 {
				if (*analysisRecordList5)[i].Typea == "1" {
					es5s2tmp := map[string]interface{}{"name": "客制签单时数(预计)"}
					es5s2 = append(es5s2, es5s2tmp)
					ac5s2tmp := map[string]interface{}{"name": "客制签单时数(实际)"}
					ac5s2 = append(ac5s2, ac5s2tmp)
				} else if (*analysisRecordList5)[i].Typea == "2" {
					es5s2tmp := map[string]interface{}{"name": "返工赠送时数(预计)"}
					es5s2 = append(es5s2, es5s2tmp)
					ac5s2tmp := map[string]interface{}{"name": "返工赠送时数(实际)"}
					ac5s2 = append(ac5s2, ac5s2tmp)
				} else if (*analysisRecordList5)[i].Typea == "3" {
					es5s2tmp := map[string]interface{}{"name": "问题处理时数(预计)"}
					es5s2 = append(es5s2, es5s2tmp)
					ac5s2tmp := map[string]interface{}{"name": "问题处理时数(实际)"}
					ac5s2 = append(ac5s2, ac5s2tmp)
				}
				es5s2s = append(es5s2s, es5s2)
				ac5s2s = append(ac5s2s, ac5s2)
			}
		}
		var es5s []map[string]interface{}
		for i := 0; i < len(es5s2s); i++ {
			es5 := make(map[string]interface{})
			for j := 0; j < len(es5s2s[i]); j++ {
				for key, value := range es5s2s[i][j] {
					es5[key] = value
				}
			}
			es5s = append(es5s, es5)
		}
		var ac5s []map[string]interface{}
		for i := 0; i < len(ac5s2s); i++ {
			ac5 := make(map[string]interface{})
			for j := 0; j < len(ac5s2s[i]); j++ {
				for key, value := range ac5s2s[i][j] {
					ac5[key] = value
				}
			}
			ac5s = append(ac5s, ac5)
		}
		esac := append(es5s, ac5s...)
		// 表头
		tbhm := make(map[string]string)
		for i := 0; i < len(es5s2s); i++ {
			for j := 0; j < len(es5s2s[i]); j++ {
				for key := range es5s2s[i][j] {
					if key != "name" {
						tbhm[key] = key
					}
				}
			}
		}
		var ths []map[string]interface{}
		ths = append(ths, map[string]interface{}{"field": "name", "title": ""})
		// 排序
		m, ks := common.Map_string(tbhm)
		for _, k := range ks {
			th := map[string]interface{}{"field": m[k], "title": m[k]}
			ths = append(ths, th)
		}
		common.OkWithDataC(count, gin.H{
			"ths":  ths,
			"esac": esac,
		}, c)
	} else {
		common.FailWithMsg("获取信息失败，请稍后重试", c)
	}
}

// 一定日期范围内的-按月按客户统计
func GetAnalysis6(c *gin.Context) {
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
	succ, analysisRecordList6, count := models.GetAnalysisRecordList6(handlerid, feedbackdatestart, feedbackdateend)
	if succ {
		// x轴
		ths_map := make(map[string]string)
		for i := 0; i < len(*analysisRecordList6); i++ {
			ths_map[(*analysisRecordList6)[i].Ny] = (*analysisRecordList6)[i].Ny
		}
		var ths []string
		for key := range ths_map {
			ths = append(ths, key)
		}
		sort.Strings(ths)
		// 客户信息
		cus_map := make(map[string]string)
		for i := 0; i < len(*analysisRecordList6); i++ {
			cus_map[(*analysisRecordList6)[i].Cname] = (*analysisRecordList6)[i].Cname
		}
		var cus []string
		for key := range cus_map {
			cus = append(cus, key)
		}
		// 数据
		var es6s2s [][]map[string]interface{}
		var es6s2 []map[string]interface{}
		for i := 0; i < len(*analysisRecordList6); i++ {
			if i == 0 {
				es6s2_m := map[string]interface{}{(*analysisRecordList6)[i].Ny: (*analysisRecordList6)[i].Handleestimatetime}
				es6s2 = append(es6s2, es6s2_m)
			} else {
				if (*analysisRecordList6)[i].Cname == (*analysisRecordList6)[i-1].Cname {
					es6s2_m := map[string]interface{}{(*analysisRecordList6)[i].Ny: (*analysisRecordList6)[i].Handleestimatetime}
					es6s2 = append(es6s2, es6s2_m)
				} else {
					es6s2_m := map[string]interface{}{"name": (*analysisRecordList6)[i-1].Cname}
					es6s2 = append(es6s2, es6s2_m)
					es6s2s = append(es6s2s, es6s2)

					es6s2 = []map[string]interface{}{}
					es6s2_m1 := map[string]interface{}{(*analysisRecordList6)[i].Ny: (*analysisRecordList6)[i].Handleestimatetime}
					es6s2 = append(es6s2, es6s2_m1)
				}
			}
			if i == len(*analysisRecordList6)-1 {
				es6s2_m := map[string]interface{}{"name": (*analysisRecordList6)[i].Cname}
				es6s2 = append(es6s2, es6s2_m)
				es6s2s = append(es6s2s, es6s2)
			}
		}
		var es6s_tmps []map[string]interface{}
		for i := 0; i < len(es6s2s); i++ {
			es6s_tmp := make(map[string]interface{})
			for j := 0; j < len(es6s2s[i]); j++ {
				for key, value := range es6s2s[i][j] {
					es6s_tmp[key] = value
				}
			}
			es6s_tmps = append(es6s_tmps, es6s_tmp)
		}
		var es6s []map[string]interface{}
		for i := 0; i < len(es6s_tmps); i++ {
			es6s_tmp1 := make(map[string]interface{})
			var es6s_data []float32
			for j := 0; j < len(ths); j++ {
				flag := false
				for key, value := range es6s_tmps[i] {
					if key == "name" {
						es6s_tmp1["name"] = value
						es6s_tmp1["type"] = "bar"
					} else {
						if key == ths[j] {
							es6s_data = append(es6s_data, value.(float32))
							flag = true
						}
					}
				}
				if !flag {
					es6s_data = append(es6s_data, 0)
				}
			}
			es6s_tmp1["data"] = es6s_data
			es6s = append(es6s, es6s_tmp1)
		}
		common.OkWithDataC(count, gin.H{
			"ths":  ths,
			"cus":  cus,
			"es6s": es6s,
		}, c)
	} else {
		common.FailWithMsg("获取信息失败，请稍后重试", c)
	}
}
