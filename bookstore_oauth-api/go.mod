module github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api

go 1.14

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/gin-gonic/gin v1.6.3
	github.com/go-resty/resty/v2 v2.3.0
	github.com/gocql/gocql v0.0.0-20200624222514-34081eda590e
	github.com/jarcoal/httpmock v1.0.5
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go v0.1.1
	github.com/stretchr/testify v1.5.1
)

replace github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go => ../bookstore_utils-go
