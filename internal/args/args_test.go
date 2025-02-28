/*
Handles args testing

Author: Wayne du Preez
*/
package args

import (
	"dupvirt/internal/logger"
	"strings"
	"testing"
)

type MockLogger struct {}

func (t *MockLogger) Debug(msg string, keysAndValues ...any) {}
func (t *MockLogger) Info(msg string, keysAndValues ...any) {}
func (t *MockLogger) Warn(msg string, keysAndValues ...any) {}
func (t *MockLogger) Error(msg string, keysAndValues ...any) {}

func TestValidate_ConfigPassEmptyFlagForDestinationServer_ExpectErrorEmptyFlagsPassed(t *testing.T) {
    /*
    We expect an err to be returned from validate as we are not allowed
    to pass empty strings for destinationServer nor interfaceMonitored
    */
    
    
    var log interface{logger.ILogger} = &MockLogger{}

    rawInputs := rawInputs{
        destinationServer: "",
        interfaceMonitored: "eth0",
        ntfyServer: "",
        ntfyTopic: "",
    }
    
    _, err := validate(log, rawInputs)

    if err != nil {
        if !strings.Contains(err.Error(), "empty values"){
            t.Log(err.Error())
            t.Errorf("Failed to get empty values substring in error")
        }
    } else {
       t.Errorf("Failed to catch error") 
    } 
}

func TestValidate_ConfigPassEmptyFlagForInterfaceMonitored_ExpectErrorEmptyFlagsPassed(t *testing.T) {
    /*
    We expect an err to be returned from validate as we are not allowed
    to pass empty strings for destinationServer nor interfaceMonitored
    */
    
    
    var log interface{logger.ILogger} = &MockLogger{}

    rawInputs := rawInputs{
        destinationServer: "gaming.home.arpa",
        interfaceMonitored: "",
        ntfyServer: "",
        ntfyTopic: "",
    }
    
    _, err := validate(log, rawInputs)

    if err != nil {
        if !strings.Contains(err.Error(), "empty values"){
            t.Log(err.Error())
            t.Errorf("Failed to get empty values substring in error")
        }
    } else {
       t.Errorf("Failed to catch error") 
    } 
}