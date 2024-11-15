package go_upstream_auth

import (
	"github.com/Kong/go-pdk"
)

func AddApiKey(kong *pdk.PDK, apiKey string, apiKeyCustomHeader string) error {
	_ = kong.Log.Debug("go-upstream-auth: AddApiKey")
	defer func() { _ = kong.Log.Debug("go-upstream-auth: AddApiKey complete") }()

	// Check if the api key is empty
	if apiKey == "" {
		_ = kong.Log.Warn("go-upstream-auth: Api key is empty")
		return nil
	}

	// Add the api key header
	if apiKeyCustomHeader == "" {
		apiKeyCustomHeader = "X-Api-Key"
	}

	err := kong.ServiceRequest.SetHeader(apiKeyCustomHeader, apiKey)
	if err != nil {
		_ = kong.Log.Err("go-upstream-auth: Error setting header: ", err)
		return err
	}

	return nil
}
