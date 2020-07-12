module github.com/olegnalivajev/learning_go_microservices/bookstore_items-api

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-go v0.0.0-00010101000000-000000000000
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go v0.1.1
)

replace (
	github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-go => ../bookstore_oauth-go
	github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go => ../bookstore_utils-go
)
