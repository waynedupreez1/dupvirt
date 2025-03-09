package pcapwrap

import (
	"fmt"
	"testing"
)

type mLogger struct{}

func (t *mLogger) Debug(msg string, keysAndValues ...any) {}
func (t *mLogger) Info(msg string, keysAndValues ...any)  {}
func (t *mLogger) Warn(msg string, keysAndValues ...any)  {}
func (t *mLogger) Error(msg string, keysAndValues ...any) {}

//This test is skipped as its testing the actual hardware on the machine
func TestDeviceExist_Config_Expect(t *testing.T) {

	//Setup
	logger := new(mLogger)
    pcap := New(logger)

	//Command
	exist, _ := pcap.DeviceExist("wlp1s0")

	t.Log(exist)

	//Test
	t.SkipNow()
	//t.Fail()
}

func Test_Config_Expect(t *testing.T) {

	logger := new(mLogger)
    
    pcap := New(logger)

	channel := pcap.GetPackets("wlp1s0", "net 192.168.25.1")

	for data := range channel{
		fmt.Println(data)
	}

	defer pcap.Close()

	t.Fail()
}