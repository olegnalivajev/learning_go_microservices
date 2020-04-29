package users

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors"
	"strings"
)

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestErr("Invalid email address.")
	}
	return nil
}