package user

import (
	log "github.com/sirupsen/logrus"
	. "tianzhuo-manage/handler"
	"tianzhuo-manage/model"
	"tianzhuo-manage/pkg/auth"
	"tianzhuo-manage/pkg/errno"
	"tianzhuo-manage/pkg/token"

	"github.com/gin-gonic/gin"
)

type Request struct {
	user    model.UserModel
	captcha string
}

// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	//验证验证码
	//if isCaptcha := captcha.Verify(c, u.Captcha); isCaptcha != true {
	//	SendResponse(c, errno.ErrCaptcha, nil)
	//	return
	//}

	// Get the user information by the login username.
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	rs := u.Cache(c)
	log.Info(rs.Val())
	SendResponse(c, nil, model.Token{Token: t})
}
