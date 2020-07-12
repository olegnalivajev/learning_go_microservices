module github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-go

go 1.14

require (
	github.com/go-resty/resty/v2 v2.3.0
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go v0.1.1
)

replace github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go => ../bookstore_utils-go
