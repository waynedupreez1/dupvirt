package common

import (
	"dupvirt/internal/logger"
	"testing"
)

type MockLogger struct{}

func (t *MockLogger) Debug(msg string, keysAndValues ...any) {}
func (t *MockLogger) Info(msg string, keysAndValues ...any)  {}
func (t *MockLogger) Warn(msg string, keysAndValues ...any)  {}
func (t *MockLogger) Error(msg string, keysAndValues ...any) {}

func Test_Config_Expect(t *testing.T) {

	var log interface{ logger.ILogger } = &MockLogger{}
	s := pass

	sendNtfy("http://cloud.home.arpa:8050", "Tests", "This is a test", "dubvirt", s, log)
}
