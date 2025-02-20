# Introduction
This project does the following:

1. Switch on a virt VM based on arp messages sent to the host
2. Switch off the virt VM once traffic to it dies down

# Whats the point
Well I have a windows game streaming machine on my network which I want to switch on and off
automatically based on usage.

# How to use

dubvirt -d <-DESTINATION SERVER IP-> -i <-INTERFACE TO MONITOR-> -n <-NOTIFY SERVER IP->

<-DESTINATION SERVER IP-> = Traffic to the Server we want to monitor, this is the server you will switch on and off
<-INTERFACE TO MONITOR-> = The interface we will monitor, we use libpcap to monitor this interface on the machine that hosts the vm
<-NOTIFY SERVER IP-> = This is the server to send notifications to, leave blank to disable