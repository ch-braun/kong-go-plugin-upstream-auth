package main

import (
	"crypto/tls"
	"github.com/Kong/go-pdk/server"
	goupstreamauth "kong-go-plugin-upstream-auth/go-upstream-auth"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const PluginVersion = "0.0.1"

var PluginPriority = 755

func main() {

	// If env var is set to false, skip TLS verification
	if strings.ToLower(os.Getenv("KONG_PLUGIN_CONFIG_GO_JWT_KC_SKIP_TLS_VERIFY")) == "true" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// If env var for priority is set, use it
	if os.Getenv("KONG_PLUGIN_CONFIG_GO_JWT_KC_PRIORITY") != "" {
		prio, err := strconv.Atoi(os.Getenv("KONG_PLUGIN_CONFIG_GO_JWT_KC_PRIORITY"))
		if err != nil {
			log.Fatalf("Error converting priority to int: %s", err)
		}
		PluginPriority = prio
	}

	_ = server.StartServer(goupstreamauth.New, PluginVersion, PluginPriority)

}
