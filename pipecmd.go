package main

import (
	//"os"
	"os/exec"
)

func pipe_commands(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].Output()
	if err != nil {
		return nil, err
	}
	return final, nil
}

func main() {
	/*
		var dirs []string
		if len(os.Args) > 1 {
			dirs = os.Args[1:]
		} else {
			dirs = []string{"."}
		}*/
	//for _, _ = range dirs {
	a1 := exec.Command("cat", "/proc/cpuinfo")
	a2 := exec.Command("grep", "-i", "Model")

	output, err := pipe_commands(a1, a2)
	if err != nil {
		println(err)
	} else {
		print(string(output))
	}
	//}
}
