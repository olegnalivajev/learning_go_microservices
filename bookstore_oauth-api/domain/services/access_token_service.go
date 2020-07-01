package services

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/model"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"strings"
)

type Repository interface {
	GetById(string) (*model.AccessToken, *errors_utils.RestErr)
	Create(model.AccessToken) *errors_utils.RestErr
	UpdateExpires(model.AccessToken) *errors_utils.RestErr
}

type AccessTokenService interface {
	GetById(string) (*model.AccessToken, *errors_utils.RestErr)
	Create(model.AccessToken) *errors_utils.RestErr
	UpdateExpires(model.AccessToken) *errors_utils.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) AccessTokenService {
	return &service{
		repository: repo,
	}
}
func (s *service) GetById(accessTokenId string) (*model.AccessToken, *errors_utils.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors_utils.NewBadRequestErr("invalid access token")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at model.AccessToken) *errors_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpires(at model.AccessToken) *errors_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpires(at)
}