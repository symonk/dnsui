package dns

import (
	"net"
)

type DnsChecker struct {
}

func New() *DnsChecker {
	return &DnsChecker{}
}

// GetIpOfHost returns the first ipv4 or ipv6 address returned
// when resolving the domain through the local resolver.
func (d *DnsChecker) GetIpOfHost(host string) string {
	addr, err := net.LookupIP(host)
	if err != nil {
		return ""
	}
	return addr[0].String()
}
