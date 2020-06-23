package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"tianzhuo-manage/handler"
	"tianzhuo-manage/pkg/errno"
	"tianzhuo-manage/util"
	"time"
)

func Upload(c *gin.Context) {
	//single file
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatalf("文件读取失败 %s", err)
	}

	//log.Info(file.Header.Values("Content-Type"))
	fileContentType := file.Header.Get("Content-Type")
	filesContentArr := strings.Split(fileContentType, "/")
	if filesContentArr[0] != "image" {
		handler.SendResponse(c, errno.NotAllowType, nil)
	}
	shortId, _ := util.GenShortId()
	filename := util.FormatDate(time.Now()) + shortId + "." + filesContentArr[1]

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "./static/image/"+filename)
	if err != nil {
		log.Panic("文件上传失败")
	}

	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}
