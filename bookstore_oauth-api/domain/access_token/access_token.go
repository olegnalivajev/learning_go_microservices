package access_token

import (
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`

	// used for password grant_type
	Username string `json:"username"`
	Password string `json:"password"`

	//used for client_credentials grant_type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Scope string `json:"scope"`
}

func (atr *AccessTokenRequest) Validate() *errors_utils.RestErr {
	switch atr.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors_utils.NewBadRequestErr("invalid grant_type parameter")
	}
	//TODO: validate params for each supported grant_type
	return nil
}

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

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) Generate() *errors_utils.RestErr {
	at.AccessToken = "pam"
	return nil
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}
