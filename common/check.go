package common

import (
	"strconv"
	"time"
)

// 检验字符串是否为时间格式
func TimeCheck(timeStr string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", timeStr)
	return err == nil
}

// 检验字符串是否为数字
func IsNumCheck(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
