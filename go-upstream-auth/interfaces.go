package go_upstream_auth

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/entities"
	"github.com/stretchr/testify/mock"
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

// --- Mocks ---

// MockPDK is a mock implementation of the PDK interface
type MockPDK struct {
	mock.Mock
}

func (m *MockPDK) Log() PDKLog {
	args := m.Called()
	return args.Get(0).(PDKLog)
}

func (m *MockPDK) Client() PDKClient {
	args := m.Called()
	return args.Get(0).(PDKClient)
}

func (m *MockPDK) Router() PDKRouter {
	args := m.Called()
	return args.Get(0).(PDKRouter)
}

func (m *MockPDK) ServiceRequest() PDKServiceRequest {
	args := m.Called()
	return args.Get(0).(PDKServiceRequest)
}

func (m *MockPDK) Response() PDKResponse {
	args := m.Called()
	return args.Get(0).(PDKResponse)
}

// MockLog is a mock implementation of the Log interface
type MockLog struct {
	mock.Mock
}

func (m *MockLog) Debug(args ...interface{}) error {
	m.Called(args)
	return nil
}

func (m *MockLog) Warn(args ...interface{}) error {
	m.Called(args)
	return nil
}

func (m *MockLog) Err(args ...interface{}) error {
	m.Called(args)
	return nil
}

// MockClient is a mock implementation of the Client interface
type MockClient struct {
	mock.Mock
}

func (m *MockClient) GetConsumer() (entities.Consumer, error) {
	args := m.Called()
	return args.Get(0).(entities.Consumer), args.Error(1)
}

// MockRouter is a mock implementation of the Router interface
type MockRouter struct {
	mock.Mock
}

func (m *MockRouter) GetRoute() (entities.Route, error) {
	args := m.Called()
	return args.Get(0).(entities.Route), args.Error(1)
}

// MockServiceRequest is a mock implementation of the ServiceRequest interface
type MockServiceRequest struct {
	mock.Mock
}

func (m *MockServiceRequest) SetHeader(key, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// MockResponse is a mock implementation of the Response interface
type MockResponse struct {
	mock.Mock
}

func (m *MockResponse) Exit(status int, body []byte, headers map[string][]string) {
	m.Called(status, body, headers)
}
