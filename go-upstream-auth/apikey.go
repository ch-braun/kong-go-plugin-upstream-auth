package go_upstream_auth

import (
	"github.com/Kong/go-pdk"
)

func AddApiKey(kong *pdk.PDK, apiKey string, apiKeyCustomHeader string) {
	_ = kong.Log.Debug("go-upstream-auth: AddApiKey")
	defer func() { _ = kong.Log.Debug("go-upstream-auth: AddApiKey complete") }()

	// Check if the api key is empty
	if apiKey == "" {
		_ = kong.Log.Warn("go-upstream-auth: Api key is empty")
		return
	}

	// Add the api key header
	if apiKeyCustomHeader == "" {
		apiKeyCustomHeader = "X-Api-Key"
	}

	_ = kong.ServiceRequest.SetHeader(apiKeyCustomHeader, apiKey)

}
