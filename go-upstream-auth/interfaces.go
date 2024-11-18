package go_upstream_auth

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/entities"
)

/*
	Here, interfaces are defined for specific sets of functionalities that the plugin needs.
	They are implemented using the Kong PDK. For testing purposes, however, interfaces are defined
	here to allow for mocking of the PDK.
*/

// PDKLog interface
type PDKLog interface {
	Debug(args ...interface{}) error
	Err(args ...interface{}) error
	Warn(args ...interface{}) error
}

// PDKClient interface
type PDKClient interface {
	GetConsumer() (entities.Consumer, error)
}

// PDKRouter interface
type PDKRouter interface {
	GetRoute() (entities.Route, error)
}

// PDKServiceRequest interface
type PDKServiceRequest interface {
	SetHeader(key string, value string) error
}

// PDKResponse interface
type PDKResponse interface {
	Exit(status int, body []byte, headers map[string][]string)
}

// PDK interface
type PDK interface {
	Log() PDKLog
	Client() PDKClient
	Router() PDKRouter
	ServiceRequest() PDKServiceRequest
	Response() PDKResponse
}

type WrappedPDK struct {
	pdk *pdk.PDK
}

func (p *WrappedPDK) Log() PDKLog {
	return p.pdk.Log
}

func (p *WrappedPDK) Client() PDKClient {
	return p.pdk.Client
}

func (p *WrappedPDK) Router() PDKRouter {
	return p.pdk.Router
}

func (p *WrappedPDK) ServiceRequest() PDKServiceRequest {
	return p.pdk.ServiceRequest
}

func (p *WrappedPDK) Response() PDKResponse {
	return p.pdk.Response
}
