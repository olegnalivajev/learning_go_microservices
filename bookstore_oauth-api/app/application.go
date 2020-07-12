package app

import (
	"github.com/gin-gonic/gin"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/access_token"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/http"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/repository/db"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/repository/rest/users"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepo(), users.NewRepo()))

	router.GET("/oauth/access_token/:id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	_ = router.Run(":8080")
}