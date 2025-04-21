package middlewares

import (
	"fmt"
	"time"

	"github.com/DucAnh2611/ls-golang/logging"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		params := map[string]string{}
		for _, p := range c.Params {
			params[p.Key] = p.Value
		}

		query := map[string]string{}
		for key, values := range c.Request.URL.Query() {
			if len(values) > 0 {
				query[key] = values[0]
			}
		}

		var body any
		if c.Request.Body != nil {
			if err := c.ShouldBindJSON(&body); err != nil {
				body = nil
			}
		}

		logStr := ""
		if len(params) > 0 {
			logStr += fmt.Sprintf(" | Params: %+v", params)
		}
		if len(query) > 0 {
			logStr += fmt.Sprintf(" | Query: %+v", query)
		}
		if body != nil {
			logStr += fmt.Sprintf(" | Body: %+v", body)
		}

		c.Next()

		duration := time.Since(start)
		logging.Info(fmt.Sprintf("%s %s | %d | %v | %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration, logStr))
	}
}
