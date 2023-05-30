package fn

import (
	"net"
)

func IntToIP(ipvi int) net.IP {
	ip := make(net.IP, 4)
	ip[0] = byte(ipvi >> 24 & 0xFF)
	ip[1] = byte(ipvi >> 16 & 0xFF)
	ip[2] = byte(ipvi >> 8 & 0xFF)
	ip[3] = byte(ipvi & 0xFF)
	return ip
}

func StringToIP(ipvs string) net.IP {
	ip := make(net.IP, 4)
	if err := ip.UnmarshalText([]byte(ipvs)); err != nil {
		return nil
	}
	return ip
}
