package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	// "strings"
)

// =============================================================================
//	What the following code does:
//	- Gets 2 environment variables CLIENT_NET and SERVER_NET
//		(for example: 172.17.0.0/16 and 172.18.0.0/16)
//	- Identifies which NIC is attached to CLIENT_NET and SERVER_NET
//	- Enables IP forwarding
//  - Adds iptables-rules in order to setup NFQUEUE between these 2 NICs
// =============================================================================

// =============================================================================
// How to test:
//	export CLIENT_NET=192.168.2.0/24
//	export SERVER_NET=192.168.3.0/24
//	go run setup_nfqueue.go
// =============================================================================

func main() {

	client_net := os.Getenv("CLIENT_NET")
	server_net := os.Getenv("SERVER_NET")

	client_nic := ""
	server_nic := ""

	if client_net != "" && server_net != "" {
		fmt.Println("Environment variables found:")
		fmt.Printf("$CLIENT_NET: %s\n", client_net)
		fmt.Printf("$SERVER_NET: %s\n", server_net)

		//try to find interfaces attached to these networks
		client_nic = GetIfaceConnectedTo(client_net)
		server_nic = GetIfaceConnectedTo(server_net)

		if client_nic != "not found" && server_nic != "not found" {
			fmt.Printf("iface attached to client_net (%s): %s\n", client_net, client_nic)
			fmt.Printf("iface attached to server_net (%s): %s\n", server_net, server_nic)

			//
			//	setup IP forwarding
			//
			cmd := "sysctl -w net.ipv4.ip_forward=1"
			fmt.Println("exec command: ", cmd)
			out, err := exec.Command("sh", "-c", cmd).Output()
			checkErr(err)
			fmt.Print(string(out))

			//
			// add iptables rule for CLIENT_NIC
			//
			cmd = "iptables -t raw -A PREROUTING -i " + client_nic + " -j NFQUEUE --queue-num 0"
			fmt.Println("exec command: ", cmd)
			out, err = exec.Command("sh", "-c", cmd).Output()
			checkErr(err)
			fmt.Print(string(out))

			//
			// add iptables rule for SERVER_NIC
			//
			cmd = "iptables -t raw -A PREROUTING -i " + server_nic + " -j NFQUEUE --queue-num 0"
			fmt.Println("exec command: ", cmd)
			out, err = exec.Command("sh", "-c", cmd).Output()
			checkErr(err)
			fmt.Print(string(out))

		} else {
			fmt.Println("Interfaces error: couldn't find NICs attaced to the networks provided")
		}

	} else {
		fmt.Println("Environment variables error: check for null $CLIENT_NET or $SERVER_NET")
	}

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
