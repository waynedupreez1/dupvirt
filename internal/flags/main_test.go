// Handles all of the flags tests
// Author: Wayne du Preez

package flags

import (
	"strings"
	"testing"

	"github.com/waynedupreez1/duptfy/internal/logger"
)

type MockLogger struct {}

func (t *MockLogger) Debug(msg string, keysAndValues ...any) {}
func (t *MockLogger) Info(msg string, keysAndValues ...any) {}
func (t *MockLogger) Warn(msg string, keysAndValues ...any) {}
func (t *MockLogger) Error(msg string, keysAndValues ...any) {}

func TestValidate_ConfigPassEmptyFlags_ExpectErrorEmptyFlagsPassed(t *testing.T) {
    
    var log interface{logger.ILogger} = &MockLogger{}

    flags := Flags{
        rawServer: "",
        rawCommand: "",
        rawMessage: "",
        Server: nil,
        Command: "",
        Message: "",
    }
    
    _, err := flags.validate(log)
    if err != nil {
        if !strings.Contains(err.Error(), "empty flags passed"){
            t.Failed()
        }
    }
}

func TestValidate_ConfigPassNotGoodUrl_ExpectInvalidUrl(t *testing.T) {
    
    var log interface{logger.ILogger} = &MockLogger{}

    flags := Flags{
        rawServer: "notGoodUrl",
        rawCommand: "blah",
        rawMessage: "blah",
        Server: nil,
        Command: "",
        Message: "",
    }
    
    _, err := flags.validate(log)
    if err != nil {
        if !strings.Contains(err.Error(), "invalid url"){
            t.Failed()
        }
    }
}
