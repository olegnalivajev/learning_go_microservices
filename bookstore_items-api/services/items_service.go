package services

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_items-api/domain/items"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors_utils.RestErr)
	Get(string) (*items.Item, *errors_utils.RestErr)
}

type itemsService struct {}

func (s *itemsService) Create(items items.Item) (*items.Item, *errors_utils.RestErr) {
	return nil, errors_utils.NewNotImplementedError()
}
func (s *itemsService) Get(id string) (*items.Item, *errors_utils.RestErr) {
	return nil, errors_utils.NewNotImplementedError()
}