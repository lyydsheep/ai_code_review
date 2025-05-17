package middleware

import (
	"bytes"
	"errors"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/common/util"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (blw *bodyLogWriter) Write(p []byte) (int, error) {
	blw.bodyBuf.Write(p)
	return blw.ResponseWriter.Write(p)
}

func StartTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("traceId")
		parentId := c.GetHeader("parentId")
		spanId := util.GenerateSpanId(c.Request.RemoteAddr)
		if traceId == "" {
			traceId = spanId
		}
		c.Set("traceId", traceId)
		c.Set("parentId", parentId)
		c.Set("spanId", spanId)
		c.Next()
	}
}

func PanicRecorder() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne.Err, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.New(c).Error("http request broken pipe", "path", c.Request.URL.Path, "error", err, "request", string(httpRequest))
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				log.New(c).Error("http request error", "path", c.Request.URL.Path,
					"error", err, "request", string(httpRequest), "stack", string(debug.Stack()))
				c.AbortWithError(http.StatusInternalServerError, err.(error))
			}
		}()
		c.Next()
	}
}

// LogAccess 记录请求和响应信息
// 请求路径、请求时间、请求方法、请求体
func LogAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 替换 writer
		blw := &bodyLogWriter{
			bodyBuf:        bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw
		reqBody := ""
		if !strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
			// 获取请求体
			// 底层调用 io.ReadAll()函数
			bodyBytes, err := c.GetRawData()
			if err != nil {
				log.New(c).Error("get request body error", "Error:", err)
				c.AbortWithStatusJSON(500, gin.H{
					"code":    500,
					"message": "get request body error",
				})
			}
			reqBody = string(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		// 多次读取，不受关闭影响
		start := time.Now()
		accessLog(c, "access log", "access_start",
			reqBody, time.Since(start), "")
		defer func() {
			outData := ""
			if c.Writer.Size() < 10*1024 {
				outData = blw.bodyBuf.String()
			}
			// 记录请求和响应信息
			accessLog(c, "access log", "access_end",
				reqBody, time.Since(start), outData)
		}()
		c.Next()
		return
	}
}

// 请求和响应的日志记录都复用这个函数
func accessLog(c *gin.Context, msg, accessType, body string, costTime time.Duration, outData string) {
	req := c.Request
	log.New(c).
		Info(msg,
			"accessType", accessType,
			"clientIp", c.ClientIP(),
			// TODO 增加 token
			//"token",
			"method", req.Method,
			"path", req.URL.Path,
			"query", req.URL.Query(),
			"reqBody", body,
			"respBody", outData,
			"costTime(ms)", costTime/time.Millisecond)
}
