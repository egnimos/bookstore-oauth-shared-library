package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/egnimos/bookstore-oauth-shared-library/oauth/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

//SET the constant public, client-id, caller-id, access-token
const (
	headerXPublic   = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXCallerId = "X-Caller-Id"

	paramAccessToken = "access_token"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 200 * time.Millisecond,
	}
)

//accessToken
type accessToken struct {
	Id       string `json:"id"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"`
}

// IsPublic : check whether the request is public or private
func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}

	return request.Header.Get(headerXPublic) == "true"
}

//GetCallerId : this method returns the caller id
func GetCallerId(request *http.Request) int64 {
	if request == nil {
		return 0
	}

	//convert the callerID into the integer64
	callerId, err := strconv.ParseInt(request.Header.Get(headerXCallerId), 10, 64)
	if err != nil {
		return 0
	}

	return callerId
}

//GetClientId : this method returns the client id
func GetClientId(request *http.Request) int64 {
	if request == nil {
		return 0 //off or false
	}

	clientId, err := strconv.ParseInt(request.Header.Get(headerXClientId), 10, 64)
	if err != nil {
		return 0 //off or false
	}

	return clientId
}

//cleanRequest : clean or remove the header request
func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}

	//remove all the header variables
	request.Header.Del(headerXClientId)
	request.Header.Del(headerXCallerId)
}

//AuthenticateRequest : this method clean the request header and authenticate the RESTAPI based on the accesstoken
func AuthenticateRequest(request *http.Request) *rest_errors.RestErr {
	if request == nil {
		return nil
	}

	//remove all the headers file
	cleanRequest(request)

	//get the access token
	accessTokenId := strings.TrimSpace(request.URL.Query().Get(paramAccessToken))
	if accessTokenId == "" {
		return nil
	}

	//get the accessToken by passing the accessTokenId
	at, tokenErr := getAccessToken(accessTokenId)
	if tokenErr != nil {
		if tokenErr.Status == http.StatusNotFound {
			return nil
		}
		return tokenErr
	}
	//add the header to the given request
	request.Header.Add(headerXClientId, fmt.Sprintf("%v", at.ClientID))
	request.Header.Add(headerXCallerId, fmt.Sprintf("%v", at.UserID))
	return nil
}

//GetAccessToken : this returns the access token with the given request id
func getAccessToken(accessTokenId string) (*accessToken, *rest_errors.RestErr) {
	response := oauthRestClient.Get(fmt.Sprintf("/oauth/access_token/%s", accessTokenId))
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying to get the access token")
	}

	if response.StatusCode > 299 {
		restErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("cannot marshal the response error")
		}
		return nil, restErr
	}
	//if there is no error then return the accesstoken
	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal the response data")
	}

	return &at, nil
}
