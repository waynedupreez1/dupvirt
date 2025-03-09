/*Package netmonitor
This contains some common functions

Author: Wayne du Preez
*/
package netmonitor

import (
    "fmt"
    "dupvirt/internal/logger" 

    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

type network struct {
    logger logger.ILogger
    interfaceName string
}

//The INetwork interface
type INetwork interface {
    findAllDevs() (ifs []pcap.Interface, err error)
}

func (n *network) findAllDevs() (ifs []pcap.Interface, err error){
    interfaces, err := pcap.FindAllDevs()

    return interfaces, err
}

func (n *network) checkInterfaceExist() bool{

    deviceList, err := n.findAllDevs()

    if err != nil {
        n.logger.Error("checkinterfaceExist: error with %w", err)
        return false
    }

    for _, device := range deviceList {
        if device.Name == n.interfaceName {
            n.logger.Info("Interface: %s Found", n.interfaceName)
            return true
        }
    }

    n.logger.Info("Interface: %s Not Found", n.interfaceName)
    return false
}

// New creates a brand network
func New(logger logger.ILogger, interfaceName string) INetwork {
    net := network {
        logger: logger,
        interfaceName: interfaceName,
    }

    return &net
}

func NetMonitor(){

    devs, err := pcap.FindAllDevs()

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(devs)

    if handle, err := pcap.OpenLive("br1", 1600, true, pcap.BlockForever); err != nil {
        panic(err)
    } else if err := handle.SetBPFFilter("port 80"); err != nil {  // optional
        panic(err)
    } else {
        packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
        for packet := range packetSource.Packets() {
          fmt.Println(packet.String()) // Do something with a packet here.
        }
    }
}
