# Introduction
This project does the following:

1. Switch on a virt VM based on amount of traffic sent to the device
2. Switch off the virt VM once traffic to it dies down

# Whats the point
I have a windows, game streaming machine on my network which I want to switch on and off
automatically based on usage.

# Requirements
You need libpcap installed on you box, the library used by tcpdump

# How to use

```
dubvirt -d <-DESTINATION SERVER IP-> -i <-INTERFACE TO MONITOR-> -n <-NOTIFY SERVER IP->

<-DESTINATION SERVER IP-> = Traffic to the Server we want to monitor, this is the server you will switch on and off
<-INTERFACE TO MONITOR-> = The interface we will monitor, we use libpcap to monitor this interface on the machine that hosts the vm
<-NOTIFY SERVER HOSTNAME/IP-> = This is the server to send notifications to, leave blank to disable
<-NOTIFY SERVER TOPIC-> = The topic to send the nnotification to, leave blank to disable
```

# Tasks
## Test
Runs All Tests
```
go test ./...
```

## Build
Build Application add to output folder
```
mkdir -p output
go build -o output/dubvirt ./cmd/dupvirt
```