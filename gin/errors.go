package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error represents the Error object structure
type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// ShouldAbortWithError takes a status and array of errors and will abort if any errors are non-nil.
// Bool return value indicates whether or not to short-circuit
func ShouldAbortWithError(c *gin.Context, status int, message string, errs ...error) bool {
	var errors []Error
	logger := GetLogger(c)

	for _, err := range errs {
		if err != nil {
			logger.WithError(err).WithField("status", status).Error(message)
			errors = append(errors, Error{
				Status: status,
				Title:  http.StatusText(status),
				Detail: message,
			})
		}
	}

	if len(errors) > 0 {
		c.AbortWithStatusJSON(status, gin.H{
			"errors": errors,
		})
		return true
	}

	c.Next()
	return false
}
