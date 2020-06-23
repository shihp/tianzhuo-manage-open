package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/willf/pad"
	_ "github.com/willf/pad"
	"io/ioutil"
	"tianzhuo-manage/handler"
	"tianzhuo-manage/pkg/errno"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		// Skip for the health check requests.
		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// The basic informations.
		method := c.Request.Method
		ip := c.ClientIP()

		//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		// get code and message
		var response handler.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Error(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}

		//noLogBody := [3]string{"/login","/upload","/v1/user/import"}
		noLogBodyMap := map[string]bool{"/login": true, "/upload": true, "/v1/user/import": true}

		//reg := regexp.MustCompile("(^/login)")
		log.Info(path)
		log.Info(noLogBodyMap[path])

		if noLogBodyMap[path] == true {
			log.Infof("%s | %s | %s %s | %s |{code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, "***", code, message)
			return
		}
		log.WithFields(log.Fields{"env": viper.GetString("env"), "latency": latency.Milliseconds()}).Infof("%s | %s |%s | %s | %s %s | %s |{code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, bodyBytes, code, message)

	}
}
