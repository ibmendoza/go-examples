package main

import (
	"errors"
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"net"
	"os/exec"
	"runtime"
)

type cmdEvent struct{}

var shell, flag = getShellAndFlag()

func runCmd(args string) []byte {
	cmd := exec.Command(shell, flag, args)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return []byte("error")
	} else {
		s := string(output)
		s = ipaddr + ": " + s
		return []byte(s)
	}
}

func getShellAndFlag() (string, string) {
	if runtime.GOOS == "windows" {
		return "cmd", "/C"
	} else {
		return "/bin/sh", "-c"
	}
}

func (c cmdEvent) HandleEvent(event []byte) {
	output := runCmd(string(event))
	conn.Publish("reply", output)
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

var conn *iris.Connection
var ipaddr string

func main() {
	topicLimits := &iris.TopicLimits{
		EventThreads: runtime.NumCPU(),
		EventMemory:  64 * 1024 * 1024,
	}

	var err error
	conn, err = iris.Connect(55555)

	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	//get ipaddr
	ipaddr, err = externalIP()
	if err != nil {
		ipaddr = "IPADDR"
	}

	var cmdHandler = new(cmdEvent)

	conn.Subscribe("cmd", cmdHandler, topicLimits)

	defer conn.Close()

	fmt.Scanln()
}
