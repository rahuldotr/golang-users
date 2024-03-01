package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		end := time.Now()
		latency := end.Sub(start)

		log.Printf("Method: [%s] | Path:  %s | Remote Addr:  %s | Latency: %v \n",
			context.Request.Method,
			context.Request.URL.Path,
			context.Request.RemoteAddr,
			latency,
		)
	}
}
