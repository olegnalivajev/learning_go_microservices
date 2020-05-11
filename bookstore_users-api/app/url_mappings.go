package app

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/controllers/ping"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:id", users.Get)
	router.PUT("/users/:id", users.Update)
	router.PATCH("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)

	// internal
	router.GET("/internal/users/search", users.Search)
}
