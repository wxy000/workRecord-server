package check

import (
	"server/common"
	"server/models"
	"strings"
)

// 检查客户编号
func ChkCustomer(customerid string) string {
	_, _, err := models.GetCustomerInfoByCustomerid(customerid)
	if err != nil {
		return "客户编号：" + err.Error() + "；"
	}
	return ""
}

// 检查作业编号
func ChkNub(nub string) string {
	if nub == "" {
		return "作业编号：不得为空；"
	}
	return ""
}

// 检查问题标题
func ChkTitle(title string) string {
	if title == "" {
		return "问题标题：不得为空；"
	}
	return ""
}

// 检查问题描述
func ChkContent(content string) string {
	if content == "" {
		return "问题描述：不得为空；"
	}
	return ""
}

// 检查人员
func ChkUser(username string) string {
	_, _, err := models.GetUserInfoByUsername(username)
	if err != nil {
		return err.Error() + "；"
	}
	return ""
}

// 检查时间
func ChkTime(time string) string {
	if time == "" {
		return "不得为空；"
	} else {
		succ := common.TimeCheck(time)
		if !succ {
			return "格式或其他错误；"
		}
	}
	return ""
}

// 检查产品
func ChkProduct(productid string) string {
	if productid == "" {
		return "产品：不得为空；"
	} else {
		tmp := strings.Split(productid, ".")
		_, _, err := models.GetProductInfoByProductid(tmp[0])
		if err != nil {
			return "产品：" + err.Error() + "；"
		}
	}
	return ""
}

// 检查是否紧急
func ChkUrgent(urgent string) string {
	if urgent == "" {
		return "是否紧急：不得为空；"
	} else {
		if (urgent != "Y") && (urgent != "N") {
			return "是否紧急：传入值必须是Y或N；"
		}
	}
	return ""
}

// 检查问题分类
func ChkIssuemain(mainid string) (string, string) {
	if mainid == "" {
		return "问题分类：不得为空；", ""
	} else {
		tmp := strings.Split(mainid, ".")
		_, _, err := models.GetIssuemainInfoByMainid(tmp[0])
		if err != nil {
			return "问题分类：" + err.Error() + "；", ""
		} else {
			return "", tmp[0]
		}
	}
}

// 检查问题类型
func ChkIssuedetail(mainid string, detailid string) string {
	if mainid == "" {
		return "问题类型：请先处理‘问题分类’的错误；"
	} else {
		if detailid == "" {
			return "问题类型：不得为空；"
		} else {
			tmp := strings.Split(detailid, ".")
			_, issuedetailtmp, err := models.GetIssuedetailInfoByDetailid(tmp[0])
			if err != nil {
				return "问题类型：" + err.Error() + "；"
			} else {
				if issuedetailtmp.Mainid != mainid {
					return "问题类型：此问题类型不是问题分类" + mainid + "的子项；"
				}
			}
		}
	}
	return ""
}

// 检查时长
func ChkTimelong(nub string) string {
	if nub == "" {
		return "不得为空；"
	} else {
		succ := common.IsNumCheck(nub)
		if !succ {
			return "数字格式有误；"
		}
	}
	return ""
}

// 检查处理回复
func ChkReply(reply string) string {
	if reply == "" {
		return "处理回复：不得为空；"
	}
	return ""
}

// 检查案件状态
func ChkCasestatus(casestatus string) string {
	if casestatus == "" {
		return "案件状态：不得为空；"
	} else {
		tmp := strings.Split(casestatus, ".")
		if (tmp[0] != "1") && (tmp[0] != "2") {
			return "案件状态：传入值必须是1.未结案或2.已结案；"
		}
	}
	return ""
}

// 检查是否现场处理
func ChkOnsite(onsite string) string {
	if onsite == "" {
		return "是否现场处理：不得为空；"
	} else {
		tmp := strings.Split(onsite, ".")
		if (tmp[0] != "1") && (tmp[0] != "2") {
			return "是否现场处理：传入值必须是1.是或2.否；"
		}
	}
	return ""
}
