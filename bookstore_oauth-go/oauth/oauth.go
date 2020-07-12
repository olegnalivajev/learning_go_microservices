package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	headerXPublic   = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXCallerId = "X-Caller-Id"

	paramAccessToken = "access_token"
)

var (
	restClient = resty.New().
		SetHostURL("http://localhost:8080").
		SetTimeout(2 * time.Second)
)

type accessToken struct {
	Id       string `json:"id"`
	UserId   int64  `json:"user_id"`
	ClientId int64  `json:"client_id"`
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func AuthenticateRequest(request *http.Request) *errors_utils.RestErr {
	if request == nil {
		return nil
	}

	cleanRequest(request)

	accessTokenId := strings.TrimSpace(request.URL.Query().Get(paramAccessToken))
	if accessTokenId == "" {
		return nil
	}

	at, err := getAccessToken(accessTokenId)
	if err != nil {
		if err.Status == http.StatusNotFound {
			return nil
		}
		return err
	}

	request.Header.Add(headerXCallerId, fmt.Sprintf("%v", at.UserId))
	request.Header.Add(headerXClientId, fmt.Sprintf("%v", at.ClientId))

	return nil
}

func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}
	request.Header.Del(headerXClientId)
	request.Header.Del(headerXCallerId)
}

func GetCallerId(r *http.Request) int64 {
	if r == nil {
		return 0
	}
	callerId, err := strconv.ParseInt(r.Header.Get(headerXCallerId), 10, 64)
	if err != nil {
		return 0
	}
	return callerId
}

func GetClientId(r *http.Request) int64 {
	if r == nil {
		return 0
	}
	clientId, err := strconv.ParseInt(r.Header.Get(headerXClientId), 10, 64)
	if err != nil {
		return 0
	}
	return clientId
}


func getAccessToken(accessTokenId string) (*accessToken, *errors_utils.RestErr) {
	response, err := restClient.R().Get(fmt.Sprintf("/oauth/access_token/%s", accessTokenId))
	if err != nil {
		return nil, errors_utils.NewBadRequestErr("Could not get access token. Time out.")
	}

	if response.StatusCode() > 299 {
		// since the response comes from our service, the error interface is common
		// so the unmarshall should succeed, if not throw an error
		var restErr errors_utils.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, errors_utils.NewInternalServerError("invalid error interface when trying to get access token", err)
		}
		return nil, &restErr
	}

	// couldn't unmarshall User, something's wrong
	var at accessToken
	if err := json.Unmarshal(response.Body(), &at); err != nil {
		return nil, errors_utils.NewInternalServerError("error trying to unmarshall access token", err)
	}

	return &at, nil
}
