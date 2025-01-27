package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(c *gin.Context) {
	c.Next() // Esegue gli handler successivi

	// Se sono presenti errori nel contesto
	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			// Log dell'errore usando Logrus
			logrus.WithFields(logrus.Fields{
				"error":  err.Err,
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
			}).Error("An error occurred")
		}

		// Restituisce una risposta JSON con il messaggio dell'errore
		c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		c.Abort()
	}
}
