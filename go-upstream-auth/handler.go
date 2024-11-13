package go_upstream_auth

import "github.com/Kong/go-pdk"

func (conf Config) Access(kong *pdk.PDK) {
	_ = kong.Log.Debug("go-upstream-auth: Access")
}
