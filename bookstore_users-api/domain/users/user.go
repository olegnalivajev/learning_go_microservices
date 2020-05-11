package users

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"strings"
)

type Status int

const (
	Pending Status = iota
	Active
)

func (status Status) String() string {
	return [...]string{"Pending", "Active"}[status]
}

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors_utils.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Status = strings.TrimSpace(user.Status)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors_utils.NewBadRequestErr("Invalid email address.")
	}
	if user.Password == "" {
		return errors_utils.NewBadRequestErr("Please provide password")
	}
	return nil
}
