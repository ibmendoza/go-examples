package main

import (
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
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
		return output
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

var conn *iris.Connection

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

	var cmdHandler = new(cmdEvent)

	conn.Subscribe("cmd", cmdHandler, topicLimits)

	defer conn.Close()

	fmt.Scanln()
}
