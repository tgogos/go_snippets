# find_nic_from_addr.go

```golang
//
//  Provided a network address e.g. 172.17.0.0/16
//  the following function will search for a network interface that is attached
//  to this network and will return the interface name.
//
func GetIfaceConnectedTo(netAddr string) string {...}
```

# setup_nfqueue.go

### What the code does:
- Gets 2 environment variables CLIENT_NET and SERVER_NET (for example: 172.17.0.0/16 and 172.18.0.0/16)
- Identifies which NIC is attached to CLIENT_NET and SERVER_NET
- Enables IP forwarding
- Adds iptables-rules in order to setup NFQUEUE between these 2 NICs

### How to test:
```bash
export CLIENT_NET=192.168.2.0/24
export SERVER_NET=192.168.3.0/24
go run setup_nfqueue.go
```
