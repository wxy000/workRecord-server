package utils

import (
	"errors"
	"log"
	"server/models"
	"server/utils/check"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// 导入工作记录
func ImportDataWR(c *gin.Context) (bool, error) {
	file, fileheader, err := c.Request.FormFile("file")
	if err != nil {
		return false, err
	} else {
		log.Println("开始导入文件：" + fileheader.Filename)
		f, err1 := excelize.OpenReader(file)
		if err1 != nil {
			return false, err1
		} else {
			// 获取 Sheet1 上所有单元格
			rows, err2 := f.GetRows("Sheet1")
			if err2 != nil {
				return false, err2
			} else {
				errStr := ""
				// 暂存的数组
				var records []models.Records
				var record models.Records
				for i, row := range rows {
					if i == 0 {
						log.Println("这是标题行，忽略此行")
					} else {
						errStrTmp := ""
						// 规定取多少列start
						rownew := [19]string{}
						copy(rownew[:], row)
						// 规定取多少列end
						for j, colCell := range rownew {
							// 客户编号校验
							if j == 0 {
								err3 := check.ChkCustomer(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								record.Customerid = colCell
							}
							// 作业编号校验
							if j == 1 {
								err3 := check.ChkNub(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								record.Number = colCell
							}
							// 问题标题校验
							if j == 2 {
								err3 := check.ChkTitle(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								record.Title = colCell
							}
							// 问题描述校验
							if j == 3 {
								err3 := check.ChkContent(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								record.Content = colCell
							}
							// 反馈人校验
							if j == 4 {
								err3 := check.ChkUser(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "反馈人：" + err3
								}
								record.Feedbackid = colCell
							}
							// 反馈时间校验
							if j == 5 {
								err3 := check.ChkTime(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "反馈时间：" + err3
								}
								record.Feedbackdate = colCell
							}
							// 产品校验
							if j == 6 {
								err3 := check.ChkProduct(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								productidtmp := strings.Split(colCell, ".")
								record.Productid = productidtmp[0]
							}
							// 是否紧急校验
							if j == 7 {
								err3 := check.ChkUrgent(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								record.Urgent = colCell
							}
							// 问题分类校验
							if j == 8 {
								err3, maintmp := check.ChkIssuemain(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								record.Issuemainid = maintmp
							}
							// 问题类型校验
							if j == 9 {
								err3 := check.ChkIssuedetail(record.Issuemainid, colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								detailidtmp := strings.Split(colCell, ".")
								record.Issuedetailid = detailidtmp[0]
							}
							// 处理人校验
							if j == 10 {
								err3 := check.ChkUser(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "处理人：" + err3
								}
								record.Handlerid = colCell
							}
							// 预计处理时长校验
							if j == 11 {
								err3 := check.ChkTimelong(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "预计处理时长：" + err3
								}
								handleestimatetimetmp, _ := strconv.ParseFloat(colCell, 32)
								record.Handleestimatetime = float32(handleestimatetimetmp)
							}
							// 实际处理时长校验
							if j == 17 {
								if colCell == "" {
									colCell = strconv.FormatFloat(float64(record.Handleestimatetime), 'f', 2, 32)
								}
								err3 := check.ChkTimelong(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "实际处理时长：" + err3
								}
								handleactualtimetmp, _ := strconv.ParseFloat(colCell, 32)
								record.Handleactualtime = float32(handleactualtimetmp)
							}
							// 处理回复校验
							if j == 12 {
								err3 := check.ChkReply(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								record.Handlereply = colCell
							}
							// 案件状态校验
							if j == 13 {
								err3 := check.ChkCasestatus(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								casestatustmp := strings.Split(colCell, ".")
								casestatustmp1, _ := strconv.Atoi(casestatustmp[0])
								record.Casestatus = uint(casestatustmp1)
							}
							// 是否现场处理校验
							if j == 14 {
								err3 := check.ChkOnsite(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								onsitetmp := strings.Split(colCell, ".")
								onsitetmp1, _ := strconv.Atoi(onsitetmp[0])
								record.Onsite = uint(onsitetmp1)
							}
							// 结案时间校验
							if j == 15 {
								err3 := check.ChkTime(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "结案时间：" + err3
								}
								record.Closetime = colCell
							}
							// bug人校验
							// 备注校验
							if j == 18 {
								err3 := check.ChkMark(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
								record.Mark = colCell
							}
						}
						if errStrTmp != "" {
							errStrTmp = "第" + strconv.Itoa(i+1) + "行：（" + errStrTmp + "）<br>"
						}
						errStr = errStr + errStrTmp
					}
					if i != 0 {
						records = append(records, record)
					}
				}
				if errStr != "" {
					return false, errors.New(errStr)
				} else {
					// 这里写插数据库
					result := models.BatchSaveRecords(&records)
					if result != nil {
						return false, errors.New(result.Error())
					} else {
						return true, errors.New("导入成功")
					}
				}
			}
		}
	}
}
