package app

import (
	"github.com/gin-gonic/gin"
	"github.com/johnnyaustor/go-bookstore-users-api/app/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	routes()
	logger.Info("start to application")
	router.Run(":8080")
}
