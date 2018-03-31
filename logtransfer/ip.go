package main


import (
	"net"
)

var ipArrays []string

func getLocalIP() (ips []string, err error){
	ifaces, err := net.Interfaces()
	if err != nil {
		return 
	}
	
	for _, i := range ifaces {
		addrs, errRet := i.Addrs()
		if errRet != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
					ip = v.IP
					if (ip.IsGlobalUnicast()) {
						ips = append(ips, ip.String())
					}
			}
		}
	}
	return
}