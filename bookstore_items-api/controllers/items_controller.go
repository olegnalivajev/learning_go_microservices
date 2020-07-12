package controllers

import (
	"fmt"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_items-api/domain/items"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_items-api/services"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-go/oauth"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {}


func (ic *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: return error to the caller
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: Return error json to the user
	}

	fmt.Println(result)
	//TODO: Return created item with HTTP status 201 - created
}

func (ic *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}