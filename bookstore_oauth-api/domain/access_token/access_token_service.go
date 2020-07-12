package access_token

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/users"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"strings"
)

type DbRepository interface {
	GetById(string) (*AccessToken, *errors_utils.RestErr)
	Create(AccessToken) *errors_utils.RestErr
	UpdateExpires(AccessToken) *errors_utils.RestErr
}

type RestUsersRepository interface {
	LoginUser(email string, password string) (*users.User, *errors_utils.RestErr)
}

type AccessTokenService interface {
	GetById(string) (*AccessToken, *errors_utils.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *errors_utils.RestErr)
	UpdateExpires(AccessToken) *errors_utils.RestErr
}

type service struct {
	dbRepo        DbRepository
	restUsersRepo RestUsersRepository
}

func NewService(dbRepo DbRepository, usersRepo RestUsersRepository) AccessTokenService {
	return &service{
		dbRepo:        dbRepo,
		restUsersRepo: usersRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors_utils.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors_utils.NewBadRequestErr("invalid access token")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(atr AccessTokenRequest) (*AccessToken, *errors_utils.RestErr) {
	if err := atr.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support more grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(atr.Username, atr.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token
	at := GetNewAccessToken(user.Id)
	at.Generate()

	// Save generated access token into Cassandra db
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpires(at AccessToken) *errors_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpires(at)
}
