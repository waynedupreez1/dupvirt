/*Package common
This contains some common functions

Author: Wayne du Preez
*/
package common

import (
	"bytes"
	"dupvirt/internal/logger"
	"fmt"
	"net/http"
)

const Ntfy = "gaming"

type status int

const (
    pass status = iota
    failed
)

func sendNtfy(ntfyURL, topic, message, title string, status status, logger logger.ILogger) {

    var url = ntfyURL + "/" + topic
    
    var comment *bytes.Buffer   
    pri := ""
    emoji := ""
    
    switch status {
        case pass:
            if message == "" {
                comment = bytes.NewBufferString("Succeeded")
            } else {
                comment = bytes.NewBufferString(message)
            }
            pri = "low"
            emoji = "+1"

        case failed:
            if message == "" {
                comment = bytes.NewBufferString("Failed")
            } else {
                comment = bytes.NewBufferString(message)
            }
            pri = "high"
            emoji = "warning"
    }
    
    logger.Info(fmt.Sprintf("Sent to Server: %s", url))

    req, err := http.NewRequest("POST", url, comment)
    if err != nil {
        logger.Error(fmt.Sprintf("http req failed: %s", err.Error()))
    } else {
        req.Header.Set("Title", title)
        req.Header.Set("Priority", pri)
        req.Header.Set("Tags", emoji)
        http.DefaultClient.Do(req)  
    }
}
