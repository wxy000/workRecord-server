package utils

import (
	"errors"
	"log"

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
				for i, row := range rows {
					if i == 0 {
						log.Println("这是标题行，忽略此行")
					} else {
						for _, colCell := range row {
							log.Print(colCell, "\t")
						}
					}
				}
			}
		}
		return true, errors.New("导入成功")
	}
}
