//https://github.com/go-mangos/mangos/issues/176

package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gdamore/mangos"
    "github.com/gdamore/mangos/protocol/push"
    "github.com/gdamore/mangos/transport/ipc"
    "github.com/gdamore/mangos/transport/tcp"
)

func die(format string, v ...interface{}) {
    fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
    os.Exit(1)
}

func NewPool(urls ...string) mangos.Socket {
    var sock mangos.Socket
    var err error
    if sock, err = push.NewSocket(); err != nil {
        die("can't get new push socket: %s", err.Error())
    }

    sock.AddTransport(ipc.NewTransport())
    sock.AddTransport(tcp.NewTransport())

    sock.SetOption(mangos.OptionWriteQLen, 1)
    for _, url := range urls {
        if err = sock.Dial(url); err != nil {
            die("can't dial on push socket: %s", err.Error())
        }
    }
    return sock
}

func Send(s mangos.Socket, msgs ...string) {
    for _, msg := range msgs {
        s.SetOption(mangos.OptionSendDeadline, 100*time.Millisecond)
        if err := s.Send([]byte(msg)); err != nil {
            log.Println(msg, "failed:", err)
        }
    }

    if err := s.Close(); err != nil {
        log.Println("Failed Close", err)
    }
}

func main() {
    Send(NewPool("tcp://127.0.0.1:8080"), "foo", "bar")
}
