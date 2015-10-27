//http://nathanleclaire.com/blog/2014/08/03/write-a-function-similar-to-underscore-dot-jss-debounce-in-golang/

package main

import (
    "fmt"
    "time"
)

func debounce(interval time.Duration, input chan int, f func(arg int)) {
    var (
        item int
    )
    for {
        select {
        case item = <-input:
            fmt.Println("received a send on a spammy channel - might be doing a costly operation if not for debounce")
        case <-time.After(interval):
            f(item)
        }
    }
}

func main() {
    spammyChan := make(chan int, 10)
    go debounce(300*time.Millisecond, spammyChan, func(arg int) {
        fmt.Println("*****************************")
        fmt.Println("* DOING A COSTLY OPERATION! *")
        fmt.Println("*****************************")
        fmt.Println("In case you were wondering, the value passed to this function is", arg)
        fmt.Println("We could have more args to our \"compiled\" debounced function too, if we wanted.")
    })
    for i := 0; i < 10; i++ {
        spammyChan <- i
    }
    time.Sleep(500 * time.Millisecond)
}
