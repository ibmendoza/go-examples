//http://dahernan.github.io/2015/02/04/context-and-cancellation-of-goroutines/

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "sync"
    "time"

    "golang.org/x/net/context"
)

var (
    wg sync.WaitGroup
)

// main is not changed
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    fmt.Println("Hey, I'm going to do some work")

    wg.Add(1)
    go work(ctx)
    wg.Wait()

    fmt.Println("Finished. I'm going home")

}

func work(ctx context.Context) error {
    defer wg.Done()

    tr := &http.Transport{}
    client := &http.Client{Transport: tr}

    // anonymous struct to pack and unpack data in the channel
    c := make(chan struct {
        r   *http.Response
        err error
    }, 1)

    req, _ := http.NewRequest("GET", "http://localhost:1111", nil)
    go func() {
        resp, err := client.Do(req)
        fmt.Println("Doing http request is a hard job")
        pack := struct {
            r   *http.Response
            err error
        }{resp, err}
        c <- pack
    }()

    select {
    case <-ctx.Done():
        tr.CancelRequest(req)
        <-c // Wait for client.Do
        fmt.Println("Cancel the context")
        return ctx.Err()
    case ok := <-c:
        err := ok.err
        resp := ok.r
        if err != nil {
            fmt.Println("Error ", err)
            return err
        }

        defer resp.Body.Close()
        out, _ := ioutil.ReadAll(resp.Body)
        fmt.Printf("Server Response: %s\n", out)

    }
    return nil
}
