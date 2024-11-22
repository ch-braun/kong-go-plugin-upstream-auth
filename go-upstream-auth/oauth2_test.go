package go_upstream_auth

import (
	"github.com/Kong/go-pdk/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"strings"
	"testing"
)

// Mock the http transport
type MockTransport struct {
	mock.Mock
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func prepareMockPDK() (*MockPDK, *MockRouter, *MockClient, *MockLog, *MockServiceRequest) {
	mockPDK := new(MockPDK)
	mockRouter := new(MockRouter)
	mockClient := new(MockClient)
	mockLog := new(MockLog)
	mockServiceRequest := new(MockServiceRequest)

	mockPDK.On("Router").Return(mockRouter)
	mockPDK.On("Client").Return(mockClient)
	mockPDK.On("Log").Return(mockLog)
	mockPDK.On("ServiceRequest").Return(mockServiceRequest)

	return mockPDK, mockRouter, mockClient, mockLog, mockServiceRequest
}

func TestAddOAuth2(t *testing.T) {

	// Replace http.DefaultClient with the mock transport
	mockTransport := new(MockTransport)
	http.DefaultClient.Transport = mockTransport

	mockPDK, mockRouter, mockClient, mockLog, mockServiceRequest := prepareMockPDK()
	mockLog.On("Debug", mock.Anything).Return(nil)
	mockLog.On("Info", mock.Anything).Return(nil)
	mockLog.On("Warn", mock.Anything).Return(nil)
	mockLog.On("Err", mock.Anything).Return(nil)

	route := entities.Route{Id: "route-id"}
	consumer := entities.Consumer{Id: "consumer-id"}

	t.Run("Client Credentials Grant Type", func(t *testing.T) {
		mockRouter.On("GetRoute").Return(route, nil).Once()
		mockClient.On("GetConsumer").Return(consumer, nil).Once()
		mockServiceRequest.On("SetHeader", "Authorization", "Bearer test-access-token").Return(nil).Once()

		mockTransport.On("RoundTrip", mock.Anything).Return(&http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"access_token": "test-access-token", "expires_in": 300}`))}, nil).Once()

		err := AddOAuth2(mockPDK, "https://token-endpoint", GrantTypeClientCredentials, "client-id", "client-secret", "scope", "", "")

		assert.NoError(t, err)
	})

	t.Run("Password Grant Type", func(t *testing.T) {
		mockRouter.On("GetRoute").Return(route, nil).Once()
		mockClient.On("GetConsumer").Return(consumer, nil).Once()
		mockServiceRequest.On("SetHeader", "Authorization", "Bearer test-access-token").Return(nil).Once()

		mockTransport.On("RoundTrip", mock.Anything).Return(&http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"access_token": "test-access-token", "expires_in": 300}`))}, nil).Once()

		err := AddOAuth2(mockPDK, "https://token-endpoint", GrantTypePassword, "", "", "scope", "username", "password")

		assert.NoError(t, err)
	})

	t.Run("Invalid Grant Type (Silent Error)", func(t *testing.T) {
		mockRouter.On("GetRoute").Return(route, nil).Once()
		mockClient.On("GetConsumer").Return(consumer, nil).Once()

		err := AddOAuth2(mockPDK, "https://token-endpoint", "invalid-grant-type", "", "", "scope", "username", "password")

		assert.NoError(t, err)
	})

	t.Run("Error Getting Route", func(t *testing.T) {
		mockRouter.On("GetRoute").Return(entities.Route{}, assert.AnError).Once()

		err := AddOAuth2(mockPDK, "https://token-endpoint", GrantTypeClientCredentials, "client-id", "client-secret", "scope", "", "")

		assert.Error(t, err)

	})

	t.Run("Error Getting Consumer", func(t *testing.T) {
		mockRouter.On("GetRoute").Return(route, nil).Once()
		mockClient.On("GetConsumer").Return(entities.Consumer{}, assert.AnError).Once()

		err := AddOAuth2(mockPDK, "https://token-endpoint", GrantTypeClientCredentials, "client-id", "client-secret", "scope", "", "")

		assert.Error(t, err)
	})

	t.Run("Error Setting Authorization Header", func(t *testing.T) {
		mockRouter.On("GetRoute").Return(route, nil).Once()
		mockClient.On("GetConsumer").Return(consumer, nil).Once()
		mockServiceRequest.On("SetHeader", "Authorization", "Bearer test-access-token").Return(assert.AnError).Once()

		mockTransport.On("RoundTrip", mock.Anything).Return(&http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"access_token": "test-access-token", "expires_in": 300}`))}, nil).Once()

		err := AddOAuth2(mockPDK, "https://token-endpoint", GrantTypeClientCredentials, "client-id", "client-secret", "scope", "", "")

		assert.Error(t, err)
	})

}
