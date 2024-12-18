package go_upstream_auth

import (
	"encoding/base64"
)

func AddBasicAuth(kong PDK, username string, password string) error {
	_ = kong.Log().Debug("go-upstream-auth: AddBasicAuth")
	defer func(log PDKLog, args ...interface{}) { _ = log.Debug(args) }(kong.Log(), "go-upstream-auth: AddBasicAuth complete")

	// Check if the username and password are empty
	if username == "" || password == "" {
		_ = kong.Log().Warn("go-upstream-auth: Username or password is empty")
		return nil
	}

	// Encode the username and password in base64
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	// Add the basic auth header
	err := kong.ServiceRequest().SetHeader("Authorization", "Basic "+base64Encoded)
	if err != nil {
		_ = kong.Log().Err("go-upstream-auth: Error setting header: ", err)
		return err
	}

	return nil
}
