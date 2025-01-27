package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		for _, err := range c.Errors {

			logrus.WithFields(logrus.Fields{
				"error":  err.Err,
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
			}).Error("An error occurred")
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		c.Abort()
	}
}
