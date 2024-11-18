package go_upstream_auth

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/log"
	"net/http"
)

func (conf Config) Access(kong *pdk.PDK) {
	_ = kong.Log.Debug("go-upstream-auth: Access")
	defer func(Log log.Log, args ...interface{}) { _ = Log.Debug(args) }(kong.Log, "go-upstream-auth: Access complete")

	// Determine the configured auth method
	var err error
	switch conf.AuthenticationMethod {
	case "apikey":
		// Call the apikey handler
		err = AddApiKey(kong, conf.ApiKey, conf.ApiKeyCustomHeader)
		break
	case "oauth2":
		// Call the oauth2 handler
		err = AddOAuth2(kong, conf.OAuth2TokenEndpoint, conf.OAuth2GrantType, conf.OAuth2ClientID, conf.OAuth2ClientSecret, conf.OAuth2Scope, conf.OAuth2Username, conf.OAuth2Password)
		break
	case "basic":
		// Call the basic handler
		err = AddBasicAuth(kong, conf.BasicUsername, conf.BasicPassword)
		break
	default:
		_ = kong.Log.Warn("go-upstream-auth: Invalid authentication method")
		return
	}
	if err != nil {
		_ = kong.Log.Err("go-upstream-auth: Could not authenticate: ", err)
		kong.Response.Exit(http.StatusUnauthorized, []byte("Unauthorized: "+err.Error()), make(map[string][]string))
		return
	}
}
