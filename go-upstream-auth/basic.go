package go_upstream_auth

import (
	"encoding/base64"
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/log"
)

func AddBasicAuth(kong *pdk.PDK, username string, password string) {
	_ = kong.Log.Debug("go-upstream-auth: AddBasicAuth")
	defer func(Log log.Log, args ...interface{}) { _ = Log.Debug(args) }(kong.Log, "go-upstream-auth: AddBasicAuth complete")

	// Check if the username and password are empty
	if username == "" || password == "" {
		_ = kong.Log.Warn("go-upstream-auth: Username or password is empty")
		return
	}

	// Encode the username and password in base64
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	// Add the basic auth header
	_ = kong.ServiceRequest.SetHeader("Authorization", "Basic "+base64Encoded)
}
