module github.com/olegnalivajev/learning_go_microservices/bookstore_users-api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/olegnalivajev/learning_go_microservices/bookstore_items-api v0.1.1 // indirect
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go v0.1.1
)

replace (
	github.com/olegnalivajev/learning_go_microservices/bookstore_items-api => ../bookstore_items-api
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go => ../bookstore_utils-go
)
