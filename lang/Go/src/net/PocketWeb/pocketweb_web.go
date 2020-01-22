package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
)

func pocketweb() {
	hostname, err := os.Hostname()
	checkError(err)
	ip, err := externalIP()
	checkError(err)

	print("\n")
	log.Printf("[%s] Servicing '%s' on :%s", progname, Opts.Directory, Opts.Port)
	log.Printf("[%s] Try\n\t\t\t\t  http://%s:%s\n\t\t\t\t  http://localhost:%s\n\t\t\t\t  http://%s:%s", progname, ip, Opts.Port, Opts.Port, hostname, Opts.Port)
	http.Handle("/", http.FileServer(http.Dir(Opts.Directory)))
	http.ListenAndServe(":"+Opts.Port, nil)
}

//==========================================================================
// Support functions

// https://stackoverflow.com/a/23558495/2125837
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
