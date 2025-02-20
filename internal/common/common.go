/* 
This is the entrypoint of the application
Author: Wayne du Preez
*/
package entrypoint

import (
	"fmt"
    "strings"
	"os/exec"
    "net/http"
	"dupvirt/internal/flags"
	"dupvirt/internal/logger"
)

type priority int

const (
    low priority = iota
    high
)

func sendNtfy(ntfyUrl , priority priority, errorMessage string) {

    var comment *strings.Reader    
    pri := ""
    emoji := ""
    
    switch priority {
        case low:
            pri = "low"
            comment = strings.NewReader("Succeeded")
            emoji = "+1"

        case high:
            pri = "high"
            
            if len(errorMessage) != 0 {
                comment = strings.NewReader(errorMessage)
            } else {
                comment = strings.NewReader("Failed")
            }
            emoji = "warning"
    }
    
    t.logger.Info(fmt.Sprintf("Sent to Server: %s", t.flags.Server.String()))

    req, err := http.NewRequest("POST", t.flags.Server.String(), comment)
    if err != nil {
        t.logger.Error(fmt.Sprintf("http req failed: %s", err.Error()))
    } else {
        req.Header.Set("Title", t.flags.Message)
        req.Header.Set("Priority", pri)
        req.Header.Set("Tags", emoji)
        http.DefaultClient.Do(req)  
    }
}
