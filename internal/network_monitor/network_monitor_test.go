package intmonitor

import (
	//"dupvirt/internal/logger"
	"testing"
)

type MockLogger struct{}

func (t *MockLogger) Debug(msg string, keysAndValues ...any) {}
func (t *MockLogger) Info(msg string, keysAndValues ...any)  {}
func (t *MockLogger) Warn(msg string, keysAndValues ...any)  {}
func (t *MockLogger) Error(msg string, keysAndValues ...any) {}

func Test_Config_Expect(t *testing.T) {
    /*
    Here I actually point it to my ntfy server so skip it here
    */
    NetMonitor()
    
}
