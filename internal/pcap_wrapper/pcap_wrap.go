/*
Package pcapwrap
A minimal wrapper over pcap and returns what I am interested in.

Tests for this is mostly skipped as it requires you to run on the actual hardware

Author: Wayne du Preez
*/
package pcapwrap

import (
    "time"
	"dupvirt/internal/logger"
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type pcapwrap struct {
    logger logger.ILogger
    done chan bool
    dataOut chan time.Time
    handle *pcap.Handle
    device string
}

type IPcapWrap interface {
    GetPackets(device string, filter string) (err error)
    Close()
}

func (p *pcapwrap) GetPackets(device string, filter string) (dataOut chan time.Time){
    p.logger.Info("Get Packages on Interface: %s Filtered with: %s", device, filter)

    go p.goGetPackets(device, filter)

    return p.dataOut
}

func Initialize(logger logger.ILogger, device string) (p *pcapwrap, err error){
    
    exist, err := deviceExist(logger, device)
    
    if err != nil{
       return nil, err 
    }

    if !exist{

    }
        

    pc := pcapwrap{
        logger: logger,
        done: make(chan bool),
        dataOut: make(chan time.Time),
    }

    return &pc, nil
}

func deviceExist(logger logger.ILogger, device string) (exist bool, err error){
    logger.Info("Checking if interface: %s Exist", device)
    
    ifsList, err := pcap.FindAllDevs()

    if err != nil {
        err = fmt.Errorf("deviceexist: %w", err)
        logger.Error(err.Error())
        return false, err
    }

    for _, ifs := range ifsList{
        if ifs.Name == device{
            return true, nil
        }
    }

    return false, nil
}

func (p *pcapwrap) goGetPackets(device string, filter string) (err error){
    p.logger.Info("Initialize Packet Source")

    handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)

    if err != nil {
        err = fmt.Errorf("getpackets: openlive: %w", err)
        p.logger.Error(err.Error())
        return err
    }

    defer handle.Close()   

    setFilterErr := handle.SetBPFFilter(filter)
    
    if setFilterErr != nil {
        err = fmt.Errorf("getpackets: setbpffilter: %w", err)
        p.logger.Error(err.Error())
        return err
    } 
    
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for {
		select {
		case <-p.done:
            fmt.Println("CLOSE CHANNEL")
            return
        case <-packetSource.Packets():
            currentDateTime := time.Now()

            p.dataOut <- currentDateTime
	    }
    }
}

func (p *pcapwrap) Close(){
    p.logger.Info("Close Packet Source")

    p.done<- true
}
