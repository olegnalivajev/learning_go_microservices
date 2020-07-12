// is going to interact with external APIs/providers as well as database
package services

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/domain/users"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/crypro_utils"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/date_utils"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}
type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors_utils.RestErr)
	GetUser(int64) (*users.User, *errors_utils.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors_utils.RestErr)
	DeleteUser(int64) *errors_utils.RestErr
	SearchUser(string) (users.Users, *errors_utils.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors_utils.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors_utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// populate auto fields
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Status = users.Status.String(users.Active)
	user.Password = crypro_utils.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(userId int64) (*users.User, *errors_utils.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors_utils.RestErr) {
	curr, err := s.GetUser(user.Id)
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
	} else {
		curr.FirstName = user.FirstName
		curr.LastName = user.LastName
		curr.Email = user.Email
	}
	if err := curr.Update(); err != nil {
		return nil, err
	}
	return curr, nil
}

func (s *userService) DeleteUser(id int64) *errors_utils.RestErr {
	user := users.User{Id: id}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *userService) SearchUser(status string) (users.Users, *errors_utils.RestErr) {
	dao := &users.User{Status: status}
	return dao.FindByStatus()
}

func (s *userService) LoginUser(request users.LoginRequest) (*users.User, *errors_utils.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypro_utils.GetMD5(request.Password),
		Status: users.Active.String(),
	}
	if err := dao.FindByEmailAndPass(); err != nil {
		return nil, err
	}
	return dao, nil
}
