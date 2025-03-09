package netmonitor

import (
    "github.com/google/gopacket/pcap"
	"dupvirt/internal/logger"
	"testing"
)

type mLogger struct{}

func (t *mLogger) Debug(msg string, keysAndValues ...any) {}
func (t *mLogger) Info(msg string, keysAndValues ...any)  {}
func (t *mLogger) Warn(msg string, keysAndValues ...any)  {}
func (t *mLogger) Error(msg string, keysAndValues ...any) {}

type mNetwork struct {
    logger logger.ILogger
    interfaceName string
}

func (n *mNetwork) findAllDevs() (ifs []pcap.Interface, err error){
    
    var interfList []pcap.Interface
    
    interf := pcap.Interface{
        Name: "eth0",
        Description: "",
        Flags: 1,
        Addresses: nil,
    }

    interfList = append(interfList, interf)

    return interfList, nil
}

func Test_Config_Expect(t *testing.T) {
    /*
    Here I actually point it to my ntfy server so skip it here
    */

    //Setup
    MLogger := mLogger{}
    
    MNetwork := mNetwork{
        logger: &MLogger,
        interfaceName: "eth0",
    }
    
    
    checkInterfaceExist()
    
}
