module github.com/olegnalivajev/learning_go_microservices/bookstore_users-api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.7.4
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go v0.1.1
)

replace (
	github.com/olegnalivajev/learning_go_microservices/bookstore_items-api => ../bookstore_items-api
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go => ../bookstore_utils-go
)
