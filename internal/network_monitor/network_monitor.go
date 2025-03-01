/*Package intmonitor
This contains some common functions

Author: Wayne du Preez
*/
package intmonitor

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func NetMonitor(){

    devs, err := pcap.FindAllDevs()

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(devs)

    if handle, err := pcap.OpenLive("tap0", 1600, true, pcap.BlockForever); err != nil {
        panic(err)
    } else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {  // optional
        panic(err)
    } else {
        packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
        for packet := range packetSource.Packets() {
          fmt.Println(packet.String()) // Do something with a packet here.
        }
    }
}
