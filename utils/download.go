package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

//下载文件
func DownloadFile(c *gin.Context, filePath string) {
	//打开文件
	_, err := os.Open(filePath)

	//获取文件的名称
	fileName := path.Base(filePath)

	if filePath == "" || fileName == "" || err != nil {
		log.Println("获取文件失败")
		// c.Redirect(http.StatusFound, "/404")
		// common.FailWithMsg("获取文件失败，请稍后重试", c)
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName)))
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(filePath)
}
