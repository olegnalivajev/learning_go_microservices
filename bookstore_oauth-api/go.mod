module github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gocql/gocql v0.0.0-20200624222514-34081eda590e
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go v0.1.1
	github.com/stretchr/testify v1.5.1
)

replace github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go => ../bookstore_utils-go
