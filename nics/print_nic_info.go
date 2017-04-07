package main

import (
	"fmt"
	"log"
	"net"
)

func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return
	}
	for _, i := range ifaces {
		log.Printf("%+v", i)
		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}
		log.Printf("%+v", addrs)
		// for _, a := range addrs {
		// 	log.Printf("%v %v\n", i.Name, a)
		// }
		log.Println("==============================================")
	}
}

func main() {
	localAddresses()
}
