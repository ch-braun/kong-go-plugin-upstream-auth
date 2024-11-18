package go_upstream_auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddBasicAuth(t *testing.T) {
	mockPDK := new(MockPDK)
	mockLog := new(MockLog)
	mockServiceRequest := new(MockServiceRequest)

	mockPDK.On("Log").Return(mockLog)
	mockPDK.On("ServiceRequest").Return(mockServiceRequest)

	t.Run("Username and password are empty", func(t *testing.T) {
		mockLog.On("Debug", mock.Anything).Return(nil).Twice()
		mockLog.On("Warn", mock.Anything).Return(nil).Once()

		err := AddBasicAuth(mockPDK, "", "")

		assert.NoError(t, err)
		mockLog.AssertCalled(t, "Warn", mock.Anything)
	})

	t.Run("Username and password are not empty", func(t *testing.T) {
		mockLog.On("Debug", mock.Anything).Return(nil).Twice()
		mockServiceRequest.On("SetHeader", "Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=").Return(nil).Once()

		err := AddBasicAuth(mockPDK, "username", "password")

		assert.NoError(t, err)
		mockServiceRequest.AssertCalled(t, "SetHeader", "Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
	})

	t.Run("Error setting header", func(t *testing.T) {
		mockLog.On("Debug", mock.Anything).Return(nil).Twice()
		mockLog.On("Err", mock.Anything).Return(nil).Once()
		mockServiceRequest.On("SetHeader", "Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=").Return(assert.AnError).Once()

		err := AddBasicAuth(mockPDK, "username", "password")

		assert.Error(t, err)
		mockServiceRequest.AssertCalled(t, "SetHeader", "Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
		mockLog.AssertCalled(t, "Err", mock.Anything)
	})
}
