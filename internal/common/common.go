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
    "time"
)

const Ntfy = "gaming"

type status int

const (
    pass status = iota
    failed
)

func sendNtfy(ntfyURL, topic, message, title string, status status, logger logger.ILogger) (*http.Response, error) {

    client := &http.Client{
        // set the time out
        Timeout: 5 * time.Second,
    }

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
        logger.Error(fmt.Sprintf("sendntfy: could not create request: %s\n", err))
        return nil, err
    }

    req.Header.Set("Title", title)
    req.Header.Set("Priority", pri)
    req.Header.Set("Tags", emoji)

    response, err := client.Do(req)  
    if err != nil {
        logger.Error(fmt.Sprintf("sendntfy: error making http request: %s\n", err))
        return nil, err
    }

    return response, nil
}
