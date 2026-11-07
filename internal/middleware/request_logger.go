package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func RequestLogger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		requestID, _ := c.Get(RequestIDKey)

		fields := []zap.Field{
			zap.String("request_id", stringValue(requestID)),
			zap.String("http.method", c.Request.Method),
			zap.String("http.route", c.FullPath()),
			zap.String("http.path", c.Request.URL.Path),
			zap.Int("http.status_code", c.Writer.Status()),
			zap.Int("http.response_size", c.Writer.Size()),
			zap.Duration("duration_ms", duration),
			zap.String("client.ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		}

		spanCtx := trace.SpanFromContext(
			c.Request.Context(),
		).SpanContext()

		if spanCtx.IsValid() {
			fields = append(fields,
				zap.String("trace_id", spanCtx.TraceID().String()),
				zap.String("span_id", spanCtx.SpanID().String()),
			)
		}

		entry := log.With(fields...)

		switch {
		case c.Writer.Status() >= 500:
			entry.Error("request completed")
		case c.Writer.Status() >= 400:
			entry.Warn("request completed")
		default:
			entry.Info("request completed")
		}
	}
}

func stringValue(value any) string {
	if s, ok := value.(string); ok {
		return s
	}
	return ""
}
