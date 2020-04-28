package app

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/controllers/ping"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.GetUser)
}
