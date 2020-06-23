package user

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	. "tianzhuo-manage/handler"
	"tianzhuo-manage/pkg/constvar"
	"tianzhuo-manage/pkg/errno"
	"tianzhuo-manage/service"
	"tianzhuo-manage/util"
	"time"
)

// @Summary List the users in the database
// @Description List users
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.ListRequest true "List users"
// @Success 200 {object} user.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /user [get]
func Export(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "名字")
	f.SetCellValue("Sheet1", "C1", "密码")
	f.SetCellValue("Sheet1", "D1", "创建时间")
	cols := [4]string{"A", "B", "C", "D"}
	for _, col := range cols {
		for i := uint64(0); i < count; i++ {
			b, _ := jsoniter.Marshal(infos[i])
			fmt.Println(string(b))
			fmt.Println(col + strconv.Itoa(int(i+2)))
			cname := col + strconv.Itoa(int(i+2))
			if col == "A" {
				fmt.Println(infos[i].Id)
				f.SetCellValue("Sheet1", cname, infos[i].Id)
			}
			if col == "B" {
				f.SetCellValue("Sheet1", cname, infos[i].Username)
			}
			if col == "C" {
				f.SetCellValue("Sheet1", cname, infos[i].Password)
			}
			if col == "D" {
				f.SetCellValue("Sheet1", cname, infos[i].CreatedAt)
			}
		}

	}

	shortId, _ := util.GenShortId()
	date := util.FormatDate(time.Now())
	fname := "用户信息-" + date + "-" + shortId + ".xlsx"
	//// Save xlsx file by the given path.
	path, err := util.GenSysPath(constvar.StaticExcelDir, fname)
	if err != nil {
		SendResponse(c, errno.ErrFileGenerate, "")
	}

	if err := f.SaveAs(path); err != nil {
		SendResponse(c, errno.ErrExcelGenerate, "")
	}
	SendResponse(c, nil, path)
}
