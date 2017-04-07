package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	netAddr := "127.0.0.1/8"

	found_nic := ""

	//try to find interface attached to netAddr
	found_nic = GetIfaceConnectedTo(netAddr)

	fmt.Printf("iface attached to network (%s): %s\n", netAddr, found_nic)

}

//
//  Provided a network address e.g. 172.17.0.0/16
//  the following function will search for a network interface that is attached
//  to this network and will return the interface name.
//
func GetIfaceConnectedTo(netAddr string) string {

	// get ifaces
	ifaces, err := net.Interfaces()

	if err != nil {
		log.Print(fmt.Errorf("Load interfaces: %v\n", err.Error()))
	} else {

		// iterate through ifaces
		for _, i := range ifaces {

			// get addresses
			addrs, err := i.Addrs()

			if err != nil {
				log.Print(fmt.Errorf("Load addresses: %v\n", err.Error()))
				continue
			}

			// iterate through addresses
			for _, a := range addrs {

				ifaceAddr := a.String()

				if SameNet(ifaceAddr, netAddr) {
					// return interface name
					// fmt.Println(i.Name)
					return i.Name
				}
			}
		}
	}
	return "not found"
}

func SameNet(netA, netB string) bool {
	_, ipnetA, _ := net.ParseCIDR(netA)
	ipB, _, _ := net.ParseCIDR(netB)

	// fmt.Printf("\n%+v\n%+v\n", ipnetA, ipB)

	if ipnetA.Contains(ipB) {
		return true
	}
	return false
}
