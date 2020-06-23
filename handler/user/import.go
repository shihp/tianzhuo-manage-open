package user

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"tianzhuo-manage/handler"
	"tianzhuo-manage/pkg/constvar"
	"tianzhuo-manage/util"
	"time"
)

func Import(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		log.Fatalf("文件读取失败 %s", err)
	}

	//log.Info(file.Header.Values("Content-Type"))
	fileContentType := file.Header.Get("Content-Type")
	log.Info(fileContentType)

	shortId, _ := util.GenShortId()
	fname := util.FormatDate(time.Now()) + "-" + shortId + "." + "xlsx"

	path, err := util.GenSysPath(constvar.StaticUploadExcelDir, fname)

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Panic("文件上传失败")
	}

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	handler.SendResponse(c, nil, "")
}
