package users

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/users"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/logger"
	"time"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors_utils.RestErr)
}

type usersRepository struct{}

var (
	restClient = resty.New().
		SetTimeout(3 * time.Second).
		SetHostURL("https://mainurl.com")
)

func NewRepo() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors_utils.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response, err := restClient.R().
		SetBody(request).
		Post("/users/login")

	// timeout
	if err != nil {
		logger.Error("Client timed out", err)
		return nil, errors_utils.NewInternalServerError("Invalid rest-client response when trying to login user", err)
	}

	if response.StatusCode() > 299 {
		// since the response comes from our service, the error interface is common
		// so the unmarshall should succeed, if not throw an error
		var restErr errors_utils.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, errors_utils.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, &restErr
	}

	// couldn't unmarshall User, something's wrong
	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, errors_utils.NewInternalServerError("error trying to unmarshall users response", err)
	}

	return &user, nil
}
