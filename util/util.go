package util

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
	"os"
	"strconv"
	"strings"
	"time"
)

func FormatDate(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func GenSysPath(dir string, path string) (string, error) {
	if !Exists(dir) {
		err := os.Mkdir(dir, 0777)
		if err != nil {
			return "", err
		}

		err = os.Chmod(dir, 0777)
		if err != nil {
			return "", err
		}
	}
	return dir + path, nil
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

func FormatStrToDate(dateStr string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, loc)
	if err == nil {
		unixTime := theTime.Unix() //1504082441
		return unixTime
	}
	return 0
}

func FormatAmount(amount float64) string {
	logrus.Info(amount)
	s := strconv.FormatFloat(amount, 'f', -1, 64)
	strArr := strings.Split(s, ".")
	if len(strArr) == 1 {
		return strArr[0]
	}
	s1 := strArr[0] + "." + strArr[1][0:2]
	return s1
}
