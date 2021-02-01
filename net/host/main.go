package main

import (
	"fmt"
	"net"
	"os"
)

var HostIp string
var HostName, _ = os.Hostname()

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	fmt.Println("addrs: ", addrs)

	for _, value := range addrs {
		if ipNet, ok := value.(*net.IPNet); ok &&
			!ipNet.IP.IsLoopback() && !ipNet.IP.IsUnspecified() {
			if ip := ipNet.IP.To4(); ip != nil {
				HostIp = ip.String()
				break
			}
		}
	}

	if HostIp == "" {
		panic("host ip is not exist")
	}

	if HostName == "" {
		panic("host name is not exist")
	}

	fmt.Printf("hostName: %s, hostIp: %s", HostName, HostIp)
}
