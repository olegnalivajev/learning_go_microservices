package db

import (
	"github.com/gocql/gocql"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/clients/cassandra"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/access_token"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
	Create(access_token.AccessToken) *errors_utils.RestErr
	UpdateExpires(access_token.AccessToken) *errors_utils.RestErr
}

type repository struct {
}

func NewRepo() DbRepository {
	return &repository{}
}

func (rep *repository) GetById(id string) (*access_token.AccessToken, *errors_utils.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors_utils.NewNotFoundErr("no access token found")
		}
		return nil, errors_utils.NewInternalServerError("could not get access token", err)
	}
	return &result, nil
}

func (rep *repository) Create(at access_token.AccessToken) *errors_utils.RestErr  {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors_utils.NewInternalServerError("could not create access token", err)
	}
	return nil
}

func (rep *repository) UpdateExpires(at access_token.AccessToken) *errors_utils.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors_utils.NewInternalServerError("could not update expiration date time", err)
	}
	return nil
}
