# winwmi

Sample code to get WMI metrics from a Windows system using Go.

This is really just a simple layout on how it's done. It would need to 
be taylored to capture what's need.

The inital metrics are specifc to doing a cloud migration.

# Cloud Metrics

These are the basic metrics that are needed for clound migration. 

* Currently Configured 
  * Configured CPUs 
  * Processor Speed 
  * Configured RAM (GiB) 
  * OS Type and Level

* Trend Over Time (5-minute samples minimum) 
  * CPU Used (%)  
  * CPU Used (#) 
  * CPU Demand (VMware)  
  * Active memory 
  * Consumed Memory 
  * File cache memory 
  * Storage Capacity Total 
  * Storage Capacity Used 

# Installation

There are a couple of ways to capture this. 
Both require that you have [Go](https://go.dev/dl/) installed on your 
system. 

Make sure you have [go installed](https://go.dev/dl/).

## Clone This Repo

Clone the repo and run:

```cmd
git clone https://github.com/vgcrld/winwmi.git
cd winwmi
go run main.go
```

## Install Directly

If you have [go installed](https://go.dev/dl/) and configured on your windows system simply run
the following command to create the binary.

`go install github.com/vgcrld/winwmi@HEAD`

To get the most recent commit.
