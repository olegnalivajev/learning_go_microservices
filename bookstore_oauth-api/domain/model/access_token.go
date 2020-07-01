package model

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors_utils.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors_utils.NewBadRequestErr("invalid access token")
	}
	if at.UserId <= 0 {
		return errors_utils.NewBadRequestErr("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors_utils.NewBadRequestErr("invalid client id")
	}
	if at.Expires <= 0 {
		return errors_utils.NewBadRequestErr("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}
