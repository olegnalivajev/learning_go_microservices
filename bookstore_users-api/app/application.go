package app

import (
	"github.com/gin-gonic/gin"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/logger"
)

var (
	router = gin.Default()
)

func StartApplication()  {
	mapUrls()
	logger.Info("About to start the application.")
	_ = router.Run(":8080")
}
