package middleware

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LogLayout struct {
	Code          int `json:"code"`
	CostTime      time.Duration
	ClientIp      string
	RequestMethod string
	RequestURI    string
	Header        map[string][]string
	Protocl       string
	Error         string
	Body          string
}

// Logger : midlleware logger
func Log(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		c.Next()
		latencyTime := time.Since(startTime)
		body, _ := io.ReadAll(c.Request.Body)

		layout := LogLayout{
			Code:          c.Writer.Status(),
			CostTime:      latencyTime,
			ClientIp:      c.ClientIP(),
			RequestMethod: c.Request.Method,
			RequestURI:    c.Request.RequestURI,
			Header:        c.Request.Header,
			Protocl:       c.Request.Proto,
			Body:          string(body),
		}

		// write the log
		var loggerTip string
		switch layout.Code {
			case http.StatusOK:
				loggerTip = "Request Info:"
			default:
				loggerTip = "Request Error:"
		}
		
		logger.Info(
			loggerTip,
			zap.Int("status-code", layout.Code),
			zap.Duration("cost-time", layout.CostTime),
			zap.String("client-ip", layout.ClientIp),
			zap.String("request-method", layout.RequestMethod),
			zap.String("request-uri", layout.RequestURI),
			zap.Any("header", layout.Header),
			zap.String("protocl", layout.Protocl),
		)
	}
}
