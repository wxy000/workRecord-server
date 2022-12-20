package utils

import (
	"errors"
	"log"
	"server/utils/check"
	"strconv"

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
				for i, row := range rows {
					if i == 0 {
						log.Println("这是标题行，忽略此行")
					} else {
						errStrTmp := ""
						//暂存问题分类id，用以问题类型的关联校验
						issuemainidtmp := ""
						// 规定取多少列start
						rownew := [16]string{}
						copy(rownew[:], row)
						// 规定取多少列end
						for j, colCell := range rownew {
							// 客户编号校验
							if j == 0 {
								err3 := check.ChkCustomer(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 作业编号校验
							if j == 1 {
								err3 := check.ChkNub(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 问题标题校验
							if j == 2 {
								err3 := check.ChkTitle(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 问题描述校验
							if j == 3 {
								err3 := check.ChkContent(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 反馈人校验
							if j == 4 {
								err3 := check.ChkUser(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "反馈人：" + err3
								}
							}
							// 反馈时间校验
							if j == 5 {
								err3 := check.ChkTime(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "反馈时间：" + err3
								}
							}
							// 产品校验
							if j == 6 {
								err3 := check.ChkProduct(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 是否紧急校验
							if j == 7 {
								err3 := check.ChkUrgent(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 问题分类校验
							if j == 8 {
								err3, maintmp := check.ChkIssuemain(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								} else {
									issuemainidtmp = maintmp
								}
							}
							// 问题类型校验
							if j == 9 {
								err3 := check.ChkIssuedetail(issuemainidtmp, colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 处理人校验
							if j == 10 {
								err3 := check.ChkUser(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "处理人：" + err3
								}
							}
							// 预计处理时长校验
							if j == 11 {
								err3 := check.ChkTimelong(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "预计处理时长：" + err3
								}
							}
							// 实际处理时长校验
							// 处理回复校验
							if j == 12 {
								err3 := check.ChkReply(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 案件状态校验
							if j == 13 {
								err3 := check.ChkCasestatus(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 是否现场处理校验
							if j == 14 {
								err3 := check.ChkOnsite(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + err3
								}
							}
							// 结案时间校验
							if j == 15 {
								err3 := check.ChkTime(colCell)
								if err3 != "" {
									errStrTmp = errStrTmp + "结案时间：" + err3
								}
							}
							// bug人校验
							// 备注校验
						}
						issuemainidtmp = ""
						if errStrTmp != "" {
							errStrTmp = "第" + strconv.Itoa(i+1) + "行：（" + errStrTmp + "）<br>"
						}
						errStr = errStr + errStrTmp
					}
				}
				if errStr != "" {
					return false, errors.New(errStr)
				} else {
					// 这里写插数据库
					return true, errors.New("导入成功")
				}
			}
		}
	}
}
