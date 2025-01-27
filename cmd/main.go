package main

import (
	"go-error-handler/config"
	"go-error-handler/logger"
	"go-error-handler/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger // Logger globale

func main() {
	err := config.LoadConfiguration()
	// Configura Logrus con l'ambiente
	// Cambia a "production" in base all'ambiente
	log = logger.SetupLogger(config.AppConfig.Env)
	if err != nil {
		log.WithError(err).Warn("No configuration file found, using default values")
	}
	log.Info("Starting Go Server with Gin and Logrus...")

	code := initServer()

	if code != 0 {
		log.Fatal("Failed to start server!")
	}
}

func initServer() int {
	router := gin.New()
	router.Use(gin.Recovery())          // Middleware per recuperare dai panic
	router.Use(middleware.ErrorHandler) // Middleware per la gestione degli errori

	// Definizione delle rotte
	router.GET(config.AppConfig.BasePath+"/tmpRequest", getDBFileData)
	router.POST(config.AppConfig.BasePath+"/tmpInsert", writeOnDBFile)
	router.DELETE(config.AppConfig.BasePath+"/tmpDelete", removeDataOnDBFile)

	// Avvio del server
	if err := router.Run(":" + config.AppConfig.AppPort); err != nil {
		log.WithError(err).Fatal("Failed to run server")
		return -1
	}

	log.Info("Server running on port ", config.AppConfig.AppPort)
	return 0
}

func getDBFileData(c *gin.Context) {
	log.Info("Processing getDBFileData request...")

	// Simulazione di un errore
	if true {
		err := c.Error(http.ErrBodyNotAllowed) // Registra l'errore nel contesto
		log.WithError(err).Warn("Handled a bad request")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "some data"})
}

func writeOnDBFile(c *gin.Context) {
	log.Info("Processing writeOnDBFile request...")
	c.JSON(http.StatusCreated, gin.H{"message": "Data written successfully"})
}

func removeDataOnDBFile(c *gin.Context) {
	log.Info("Processing removeDataOnDBFile request...")
	c.JSON(http.StatusOK, gin.H{"message": "Data removed successfully"})
}
