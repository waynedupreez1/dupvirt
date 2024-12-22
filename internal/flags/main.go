// Handles all of the flags returned when the cli is executed
// Author: Wayne du Preez

package flags

import (
    "os"
    "fmt"
    "flag"
    "net/url"
	"github.com/waynedupreez1/duptfy/internal/logger"
)

// Flags contains all the behaviour and flags data
type Flags struct {
    rawServer string
    rawCommand string
    rawMessage string
    Server *url.URL
    Command string
    Message string
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

    var urlString,cmd, message string
    flag.StringVar(&urlString, "s", "", "Required. ntfy endpoint URL ie. http://example.com/backup.")
    flag.StringVar(&cmd, "c", "", "Required. Bash command to run ie. 'ls -als | grep blah'.")
    flag.StringVar(&message, "m", "", "Required. Message when sent to ntfy ie. 'Local Rsnapshot backup'.")

    flag.Parse()

    flags := Flags{
        rawServer: urlString, 
        rawCommand: cmd,
        rawMessage: message,
    }

    return &flags
}

func (t *Flags)validate(logger logger.ILogger) (*Flags, error) {

    logger.Info("Validate Raw Flags")

    if t.rawServer == "" || t.rawCommand == "" || t.rawMessage == "" {
        flag.Usage()
        errMsg := fmt.Errorf("empty flags passed")
        logger.Error(errMsg.Error())
        return nil,errMsg
    }

    serverURL, err := url.Parse(t.rawServer)
    if err != nil {
        errMsg := fmt.Errorf("url parsing failed with: %s", err.Error())
        logger.Error(errMsg.Error())
        return nil,errMsg
    }

    if !serverURL.IsAbs(){
        errMsg := fmt.Errorf("invalid url: %s", serverURL)
        logger.Error(errMsg.Error())
        return nil,errMsg
    }

    logger.Info(fmt.Sprintf("ntfy Server URL: %s", serverURL))
    logger.Info(fmt.Sprintf("Command that will run: %s", t.rawCommand))
    
    t.Server = serverURL
    t.Message = t.rawMessage
    t.Command = t.rawCommand

    return t, nil
}
