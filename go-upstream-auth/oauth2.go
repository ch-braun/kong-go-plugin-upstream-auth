package go_upstream_auth

import (
	"encoding/json"
	"errors"
	"github.com/Kong/go-pdk/entities"
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	GrantTypeClientCredentials    = "client_credentials"
	GrantTypePassword             = "password"
	HTTPHeaderContentType         = "Content-Type"
	HTTPContentTypeFormUrlEncoded = "application/x-www-form-urlencoded"
)

var accessTokenCache = cache.New(5*time.Minute, 10*time.Minute)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func AddOAuth2(kong PDK, oauth2TokenEndpoint string, oauth2GrantType string, oauth2ClientID string, oauth2ClientSecret string, oAuth2Scope string, oAuth2Username string, oAuth2Password string) error {
	var accessToken string
	var err error

	route, err := kong.Router().GetRoute()
	if err != nil {
		_ = kong.Log().Err("go-upstream-auth: Could not get route: ", err)
		return err
	}

	consumer, err := kong.Client().GetConsumer()
	if err != nil {
		_ = kong.Log().Err("go-upstream-auth: Could not get consumer: ", err)
		return err
	}

	switch oauth2GrantType {
	case GrantTypeClientCredentials:
		// Call the client_credentials handler
		accessToken, err = fetchAccessTokenWithClientCredentials(&route, &consumer, oauth2TokenEndpoint, oauth2ClientID, oauth2ClientSecret, oAuth2Scope)
		break
	case GrantTypePassword:
		// Call the password handler
		accessToken, err = fetchAccessTokenWithPassword(&route, &consumer, oauth2TokenEndpoint, oAuth2Scope, oAuth2Username, oAuth2Password)
		break
	default:
		_ = kong.Log().Warn("go-upstream-auth: Invalid grant type")
		return nil
	}

	if err != nil {
		_ = kong.Log().Err("go-upstream-auth: Could not fetch access token: ", err)
		return err
	}

	err = kong.ServiceRequest().SetHeader("Authorization", "Bearer "+accessToken)
	if err != nil {
		_ = kong.Log().Err("go-upstream-auth: Could not set Authorization header: ", err)
		return err
	}

	_ = kong.Log().Debug("go-upstream-auth: Authorization header set")
	return nil
}

func prepareAccessTokenRequestBody(grantType string, username string, password string, scope string) string {
	accessTokenRequest := "grant_type=" + grantType
	if username != "" {
		accessTokenRequest += "&username=" + username
	}
	if password != "" {
		accessTokenRequest += "&password=" + password
	}
	if scope != "" {
		accessTokenRequest += "&scope=" + scope
	}
	return accessTokenRequest
}

func doAccessTokenRequest(oauth2TokenEndpoint string, accessTokenRequest string, oauth2ClientID *string, oauth2ClientSecret *string) (*AccessTokenResponse, error) {
	var accessTokenResponse AccessTokenResponse

	req, err := http.NewRequest(http.MethodPost, oauth2TokenEndpoint, strings.NewReader(accessTokenRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set(HTTPHeaderContentType, HTTPContentTypeFormUrlEncoded)

	if oauth2ClientID != nil && oauth2ClientSecret != nil {
		req.SetBasicAuth(*oauth2ClientID, *oauth2ClientSecret)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New("access token request failed with status code: " + resp.Status)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &accessTokenResponse)
	if err != nil {
		return nil, err
	}

	return &accessTokenResponse, nil
}

func fetchAccessTokenWithClientCredentials(route *entities.Route, consumer *entities.Consumer, oauth2TokenEndpoint string, oauth2ClientID string, oauth2ClientSecret string, oAuth2Scope string) (string, error) {
	cacheKey := "cc:" + oauth2TokenEndpoint + ":" + oauth2ClientID + ":" + oAuth2Scope + ":" + route.Id + ":" + consumer.Id
	cachedAccessToken, found := accessTokenCache.Get(cacheKey)
	if found {
		return cachedAccessToken.(string), nil
	}

	accessTokenRequest := prepareAccessTokenRequestBody(GrantTypeClientCredentials, "", "", oAuth2Scope)
	accessTokenResponse, err := doAccessTokenRequest(oauth2TokenEndpoint, accessTokenRequest, &oauth2ClientID, &oauth2ClientSecret)

	if err != nil {
		return "", err
	}

	if accessTokenResponse.AccessToken == "" {
		return "", errors.New("access token is empty")
	}

	accessTokenCache.Set(cacheKey, accessTokenResponse.AccessToken, time.Duration(accessTokenResponse.ExpiresIn)*time.Second)

	return accessTokenResponse.AccessToken, nil
}

func fetchAccessTokenWithPassword(route *entities.Route, consumer *entities.Consumer, oauth2TokenEndpoint string, oAuth2Scope string, oAuth2Username string, oAuth2Password string) (string, error) {
	cacheKey := "pw:" + oauth2TokenEndpoint + ":" + oAuth2Username + ":" + oAuth2Scope + ":" + route.Id + ":" + consumer.Id
	cachedAccessToken, found := accessTokenCache.Get(cacheKey)
	if found {
		return cachedAccessToken.(string), nil
	}

	accessTokenRequest := prepareAccessTokenRequestBody(GrantTypePassword, oAuth2Username, oAuth2Password, oAuth2Scope)
	accessTokenResponse, err := doAccessTokenRequest(oauth2TokenEndpoint, accessTokenRequest, nil, nil)

	if err != nil {
		return "", err
	}

	if accessTokenResponse.AccessToken == "" {
		return "", errors.New("access token is empty")
	}

	accessTokenCache.Set(cacheKey, accessTokenResponse.AccessToken, time.Duration(accessTokenResponse.ExpiresIn)*time.Second)

	return accessTokenResponse.AccessToken, nil
}
