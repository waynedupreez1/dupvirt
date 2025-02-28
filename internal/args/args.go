/*Package args
Handles all the command line variables passed to the program

Author: Wayne du Preez
*/
package args

import (
	"dupvirt/internal/logger"
	"flag"
	"fmt"
)

// These are the raw flags passed from the command line
type rawInputs struct {
    destinationServer string
    interfaceMonitored string
    ntfyServer string
    ntfyTopic string
}

// Inputs from command line after validation
type Inputs struct {
    DestinationServer string
    InterfaceMonitored string    
    NtfyServer string
    NtfyTopic string
}

// Get will retrieve and validate all command line inputs
func Get(logger logger.ILogger) *Inputs {

    logger.Info("Get Command Line Arguments Passed")

    rInputs := initialize(logger)
    vInputs, err := validate(logger, rInputs)
    
    if err != nil {
        panic("Unrecoverable error in validating passed arguments")
    }

    return vInputs
}

func initialize(logger logger.ILogger) rawInputs {

    raw := rawInputs{
        destinationServer: "", 
        interfaceMonitored: "",
        ntfyServer: "",
        ntfyTopic: "",
    }
    
    flag.StringVar(&raw.destinationServer, "d", "", "Required. Traffic to the Server we want to monitor, this is the server you will switch on and off. ie. gaming.home.arpa")
    flag.StringVar(&raw.interfaceMonitored, "i", "", "Required. The interface we will monitor, we use libpcap to monitor this interface on the machine that hosts the vm. ie eth0")
    flag.StringVar(&raw.ntfyServer, "n", "empty", "Optional. This is the ntfy server to send notifications to, leave blank to disable.")
    flag.StringVar(&raw.ntfyTopic, "t", "empty", "Optional. This is the ntfy topic to send notifications to, leave blank to disable.")

    flag.Parse()

    return raw
}

func validate(logger logger.ILogger, rawInput rawInputs) (*Inputs, error) {

    logger.Info("Validate Raw Command Line Arguments")

    if rawInput.destinationServer == "" || rawInput.interfaceMonitored == "" {
        flag.Usage()
        errMsg := fmt.Errorf("empty values, neither -d nor -i flag can be empty")
        logger.Error(errMsg.Error())
        return nil, errMsg
    }

    logger.Info(fmt.Sprintf("Destination Server URL: %s", rawInput.destinationServer))
    logger.Info(fmt.Sprintf("Monitoring Interface: %s", rawInput.interfaceMonitored))    
    logger.Info(fmt.Sprintf("Ntfy Server URL: %s",  rawInput.ntfyServer))
    logger.Info(fmt.Sprintf("Ntfy Topic: %s",  rawInput.ntfyTopic))

    validInputs := Inputs {
        DestinationServer: rawInput.destinationServer,
        InterfaceMonitored: rawInput.interfaceMonitored,
        NtfyServer: rawInput.ntfyServer,
        NtfyTopic: rawInput.ntfyTopic,
    }

    return &validInputs, nil
}
