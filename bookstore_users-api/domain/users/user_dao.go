package users

import (
	"fmt"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr  {
	result := usersDb[user.Id]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("User %d not found.", user.Id))
	}

	user.Id = result.Id
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDb[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestErr("User with this email address already exists.")
		}
		return errors.NewBadRequestErr(fmt.Sprintf("User %d already exists.", user.Id))
	}
	usersDb[user.Id] = user
	return nil
}
