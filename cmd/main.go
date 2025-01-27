package main

import (
	"go-error-handler/config"
	"go-error-handler/logger"
	"go-error-handler/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfiguration()

	log := logger.SetupLogger(config.AppConfig.Env)
	log.Info("Starting Go Server with Gin and Logrus...")

	if err := initServer(log); err != nil {
		log.WithError(err).Fatal("Failed to start server!")
	}
}

func initServer(log *logrus.Logger) error {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler)

	router.GET(config.AppConfig.BasePath+"/tmpRequest", func(c *gin.Context) {
		getDBFileData(c, log)
	})
	router.POST(config.AppConfig.BasePath+"/tmpInsert", func(c *gin.Context) {
		writeOnDBFile(c, log)
	})
	router.DELETE(config.AppConfig.BasePath+"/tmpDelete", func(c *gin.Context) {
		removeDataOnDBFile(c, log)
	})

	log.Infof("Server running on port %s", config.AppConfig.AppPort)
	return router.Run(":" + config.AppConfig.AppPort)
}

func getDBFileData(c *gin.Context, log *logrus.Logger) {
	log.Info("Processing getDBFileData request...")

	if true {
		err := c.Error(http.ErrBodyNotAllowed)
		log.WithError(err).Warn("Handled a bad request")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "some data"})
}

func writeOnDBFile(c *gin.Context, log *logrus.Logger) {
	log.Info("Processing writeOnDBFile request...")
	c.JSON(http.StatusCreated, gin.H{"message": "Data written successfully"})
}

func removeDataOnDBFile(c *gin.Context, log *logrus.Logger) {
	log.Info("Processing removeDataOnDBFile request...")
	c.JSON(http.StatusOK, gin.H{"message": "Data removed successfully"})
}
