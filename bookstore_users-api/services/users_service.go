package services

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/domain/users"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr)  {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	curr, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			curr.FirstName = user.FirstName
		}
		if user.LastName != "" {
			curr.LastName = user.LastName
		}
		if user.Email != "" {
			curr.Email = user.Email
		}
	} else{
		curr.FirstName = user.FirstName
		curr.LastName = user.LastName
		curr.Email = user.Email
	}
	if err := curr.Update(); err != nil {
		return nil, err
	}
	return curr, nil
}
