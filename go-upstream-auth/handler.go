package go_upstream_auth

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/log"
)

func (conf Config) Access(kong *pdk.PDK) {
	_ = kong.Log.Debug("go-upstream-auth: Access")
	defer func(Log log.Log, args ...interface{}) { _ = Log.Debug(args) }(kong.Log, "go-upstream-auth: Access complete")

	// Determine the configured auth method
	switch conf.AuthenticationMethod {
	case "apikey":
		// Call the apikey handler
		AddApiKey(kong, conf.ApiKey, conf.ApiKeyCustomHeader)
		break
	case "oauth2":
		// Call the oauth2 handler
		break
	case "basic":
		// Call the basic handler
		break
	default:
		_ = kong.Log.Warn("go-upstream-auth: Invalid authentication method")
		return
	}
}
