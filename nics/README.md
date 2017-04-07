# find_nic_from_addr.go

```golang
//
//  Provided a network address e.g. 172.17.0.0/16
//  the following function will search for a network interface that is attached
//  to this network and will return the interface name.
//
func GetIfaceConnectedTo(netAddr string) string {...}
```
