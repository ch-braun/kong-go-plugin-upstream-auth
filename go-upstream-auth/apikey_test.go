package go_upstream_auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddApiKey(t *testing.T) {
	mockPDK := new(MockPDK)
	mockLog := new(MockLog)
	mockServiceRequest := new(MockServiceRequest)

	mockPDK.On("Log").Return(mockLog)
	mockPDK.On("ServiceRequest").Return(mockServiceRequest)

	t.Run("ApiKey is empty", func(t *testing.T) {
		mockLog.On("Debug", mock.Anything).Return(nil).Twice()
		mockLog.On("Warn", mock.Anything).Return(nil).Once()

		err := AddApiKey(mockPDK, "", "")

		assert.NoError(t, err)
		mockLog.AssertCalled(t, "Warn", mock.Anything)
	})

	t.Run("ApiKey is not empty, custom header is empty", func(t *testing.T) {
		mockLog.On("Debug", mock.Anything).Return(nil).Twice()
		mockServiceRequest.On("SetHeader", "X-Api-Key", "test-api-key").Return(nil).Once()

		err := AddApiKey(mockPDK, "test-api-key", "")

		assert.NoError(t, err)
		mockServiceRequest.AssertCalled(t, "SetHeader", "X-Api-Key", "test-api-key")
	})

	t.Run("ApiKey is not empty, custom header is provided", func(t *testing.T) {
		mockLog.On("Debug", mock.Anything).Return(nil).Twice()
		mockServiceRequest.On("SetHeader", "Custom-Header", "test-api-key").Return(nil).Once()

		err := AddApiKey(mockPDK, "test-api-key", "Custom-Header")

		assert.NoError(t, err)
		mockServiceRequest.AssertCalled(t, "SetHeader", "Custom-Header", "test-api-key")
	})

	t.Run("Error setting header", func(t *testing.T) {
		mockLog.On("Debug", mock.Anything).Return(nil).Twice()
		mockLog.On("Err", mock.Anything).Return(nil).Once()
		mockServiceRequest.On("SetHeader", "X-Api-Key", "test-api-key").Return(assert.AnError).Once()

		err := AddApiKey(mockPDK, "test-api-key", "")

		assert.Error(t, err)
		mockServiceRequest.AssertCalled(t, "SetHeader", "X-Api-Key", "test-api-key")
		mockLog.AssertCalled(t, "Err", mock.Anything)
	})
}
