package middleware

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

var requestIDPattern = regexp.MustCompile(`^[a-zA-Z0-9._-]{1,128}$`)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := strings.TrimSpace(
			c.GetHeader("X-Request-ID"),
		)

		// Only accept a safe, bounded request ID.
		if !requestIDPattern.MatchString(requestID) {
			requestID = uuid.NewString()
		}

		c.Set(RequestIDKey, requestID)
		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}
