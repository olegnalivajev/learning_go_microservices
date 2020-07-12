package app

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/:id", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}