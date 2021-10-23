package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Log() gin.HandlerFunc {
	filePath := "log/log.log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("open log file error!", err)
	}
	// logger
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)

	// 定制化logger 组件
	logWriter, _ := retalog.New(
		filePath+"%Y-%m-%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		// 可以软连接到一个新的文件，将最新的日志软连接到最新的log
		//retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
	// 中间件
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // continue task
		stopTime := time.Since(startTime).Milliseconds()
		Cost := fmt.Sprintf("%d ms", stopTime)
		hostName, err := os.Hostname() // 获取到客户端的类型
		if err != nil {
			hostName = "unknown"
		}
		// 获取相关信息
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		UA := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"Cost":      Cost,
			"ip":        clientIp,
			"UserAgent": UA,
			"method":    method,
			"path":      path,
			"dataSize":  dataSize,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		//根据不同的状态码输出不同等级的log
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
