package users

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors"
	"strings"
)

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestErr("Invalid email address.")
	}
	return nil
}