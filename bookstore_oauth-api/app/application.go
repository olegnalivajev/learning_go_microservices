package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/clients/cassandra"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/services"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/http"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	checkCassandraAlive()

	atHandler := http.NewHandler(services.NewService(db.NewRepo()))

	router.GET("/oauth/access_token/:id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	_ = router.Run(":8080")
}

func checkCassandraAlive() {
	session, err := cassandra.GetSession()
	if err != nil {
		fmt.Println("Could not start Cassandra cluster.")
		panic(err)
	}
	session.Close()
}