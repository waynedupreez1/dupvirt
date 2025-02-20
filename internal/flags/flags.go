/* 
Package flags handles all of the flags returned when the cli is executed
Author: Wayne du Preez
*/
package flags

import (
    "os"
    "fmt"
    "flag"
    "net/url"
	"dupvirt/internal/logger"
)

// Flags contains all the behaviour and flags data
type Flags struct {
    rawDestinationServer string
    rawNtfyServer string    
    monitorInterface string
    DestinationServer *url.URL
    NtfyServer *url.URL
    MonitorInterface string    
}

// Get will retrive and validate all passed flags
func Get(logger logger.ILogger) *Flags {

    logger.Info("Get Flags")

    rflags := initialize(logger)
    vflag, err := rflags.validate(logger)
    
    if err != nil {
        os.Exit(1)
    }

    return vflag
}

func initialize(logger logger.ILogger) *Flags {

    logger.Info("Return Raw Flags")

    var rawDestinationServer, rawNtfyServer, monitorInterface  string
    flag.StringVar(&rawDestinationServer, "d", "", "Required. Traffic to the Server we want to monitor, this is the server you will switch on and off. ie. gaming.home.arpa")
    flag.StringVar(&monitorInterface, "i", "", "Required. The interface we will monitor, we use libpcap to monitor this interface on the machine that hosts the vm. ie eth0")
    flag.StringVar(&rawNtfyServer, "n", "empty", "Optional. This is the server to send notifications to, leave blank to disable.")

    flag.Parse()

    flags := Flags{
        rawDestinationServer: rawDestinationServer, 
        monitorInterface: monitorInterface,
        rawNtfyServer: rawNtfyServer,
    }

    return &flags
}

func (flags *Flags)validate(logger logger.ILogger) (*Flags, error) {

    logger.Info("Validate Raw Flags")

    if flags.rawDestinationServer == "" || flags.monitorInterface == "" {
        flag.Usage()
        errMsg := fmt.Errorf("neither -d nor -n flag can be empty")
        logger.Error(errMsg.Error())
        return nil, errMsg
    }

    destinationServer, err := url.Parse(flags.rawDestinationServer)
    if err != nil {
        errMsg := fmt.Errorf("destination server url parsing failed with: %s", err.Error())
        logger.Error(errMsg.Error())
        return nil, errMsg
    }

    //Ensure the url has a schema, meaning http/https/etc
    if !destinationServer.IsAbs(){
        errMsg := fmt.Errorf("destination server url does not have a valid schema ie. http/https: %w", destinationServer)
        logger.Error(errMsg.Error())
        return nil, errMsg
    }

    if flags.rawNtfyServer != "empty"{
        ntfyServer, err := url.Parse(flags.rawNtfyServer)
        if err != nil {
            errMsg := fmt.Errorf("ntfy server url does not have a valid schema ie. http/https: %w", err.Error())
            logger.Error(errMsg.Error())
            return nil, errMsg
        }      
    }


    logger.Info(fmt.Sprintf("ntfy Server URL: %s", serverURL))
    logger.Info(fmt.Sprintf("Command that will run: %s", t.rawCommand))
    
    t.Server = serverURL
    t.Message = t.rawMessage
    t.Command = t.rawCommand

    return t, nil
}
