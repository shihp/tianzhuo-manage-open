package user

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	. "tianzhuo-manage/handler"
	"tianzhuo-manage/model"
	"tianzhuo-manage/pkg/errno"
)

func Jmeter(c *gin.Context) {

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	//1s 5000~6000
	rs := u.Cache(c)
	log.Info(rs.Val())

	//1s 1w
	u.Mongo(c)
	SendResponse(c, nil, "success")
}
