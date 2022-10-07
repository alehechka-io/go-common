package gin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const traceID string = "traceID"

// Logger is a middleware to log  request information with logrus
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		trace := uuid.New()
		ctx.Set(traceID, trace.String())
		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)

		logger := log.WithField(traceID, trace).
			WithField("method", ctx.Request.Method).
			WithField("path", ctx.Request.URL.Path).
			WithField("status", ctx.Writer.Status()).
			WithField("ip", ctx.ClientIP()).
			WithField("user_agent", ctx.Request.UserAgent()).
			WithField("latency", latency).
			WithTime(end)

		logger.Info("[GIN]")
	}
}

// GetLogger creates log entry with traceID appended
func GetLogger(ctx *gin.Context) *log.Entry {
	trace, ok := ctx.Get(traceID)

	if ok {
		return log.WithField(traceID, trace)
	}
	return log.NewEntry(log.StandardLogger())
}
