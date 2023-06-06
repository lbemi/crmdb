package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		//start := time.Now()
		//path := c.Request.URL.Path
		//query := c.Request.URL.RawQuery
		log.Logger.Info("test---middleware")
		c.Next()

	}
}
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		str := fmt.Sprintf(" %v    status: %v, method: %v , query: %v, "+
			"IP: %v, Remote-IP: %v, USER-AGENT: %v, errors: %v, cost: %v ",
			path, c.Writer.Status(), c.Request.Method, query,
			c.ClientIP(), c.RemoteIP(), c.Request.UserAgent(), c.Errors.ByType(gin.ErrorTypePrivate).String(),
			cost)
		log.Logger.Info(str)
	}
}

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.Logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					errStr := fmt.Sprintf("[Recovery from panic]  error: %v \n %v ", err, string(debug.Stack()))
					log.Logger.Error(errStr)
				} else {
					errStr := fmt.Sprintf("[Recovery from panic]  error: %v ", err)
					log.Logger.Error(errStr)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
