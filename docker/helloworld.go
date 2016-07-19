package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip, _ := externalIP()
	fmt.Fprintf(w, "Hi there, I love %s! From: %s ", r.URL.Path[1:], ip)

	//ipaddr(w, r)

	ipaddr2(w, r)
}

//https://github.com/crosbymichael/ip-addr
func ipaddr(w http.ResponseWriter, r *http.Request) {
	i, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatal(err)
	}
	addrs, err := i.Addrs()
	if err != nil {
		log.Fatal(err)
	}
	// get the highest order ip
	if len(addrs) == 0 {
		return
	}
	for _, a := range addrs {
		n, ok := a.(*net.IPNet)
		if !ok {
			continue
		}
		//fmt.Print(n.IP)
		fmt.Fprintf(w, "From: eth0 %s ", n.IP)
		return
	}
}

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

//http://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func ipaddr2(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	for _, a := range addrs {
		//fmt.Println(a)
		fmt.Fprintf(w, "%s \n", a)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
