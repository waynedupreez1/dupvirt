// Handles all of the cli tests
// Author: Wayne du Preez

package cli

import (
    "testing"
    "github.com/waynedupreez1/duptfy/internal/flags"
    "github.com/waynedupreez1/duptfy/internal/logger"
)

type MockLogger struct {}

func (t *MockLogger) Debug(msg string, keysAndValues ...any) {}
func (t *MockLogger) Info(msg string, keysAndValues ...any) {}
func (t *MockLogger) Warn(msg string, keysAndValues ...any) {}
func (t *MockLogger) Error(msg string, keysAndValues ...any) {}

func TestCliRunCmd_ConfigCommandNotToExist_ExpectErrorToBeRaised(t *testing.T) {
    
    var log interface{logger.ILogger} = &MockLogger{}

    flags := flags.Flags{
        Server: nil,
        Command: "bf",
        Message: "",
    }    
    
    cli := New(log, &flags)

    _, err := cli.runCommand()

    if err == nil {
        t.Failed()
    }
}
