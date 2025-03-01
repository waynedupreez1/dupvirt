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

func TestSendNtfy_ConfigCorrectUrl_ExpectSendToNtfy(t *testing.T) {
    /*
    Here I actually point it to my ntfy server so skip it here
    */

    t.SkipNow()
    
    var log interface{ logger.ILogger } = &MockLogger{}
    s := pass

    res, err := sendNtfy("http://example.com", "Tests", "This is a test", "dubvirt", s, log)

        // Unwrap the error to check if it's a connection refused error
        // we Expect this else fail    
    // if err != nil {
    //     var netErr *net.OpError
    //     if errors.As(err, &netErr) {
    //         if netErr.Op == "dial" {
    //             if netErr.Err.Error() == "connect: connection refused" {
    //                 return
    //             }
    //         }
    //     }
    // }
    t.Log(res)
    t.Log(err)
}
