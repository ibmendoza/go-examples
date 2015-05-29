//Courtesy: Fatih Arslan et al

package main

import (
    "fmt"
    "github.com/fatih/set"
    "strconv"
    "sync"
)

func main() {
    var wg sync.WaitGroup // this is just for waiting until all goroutines finish

    // Initialize our thread safe Set
    s := set.New()

    // Add items concurrently (item1, item2, and so on)
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            item := "item" + strconv.Itoa(i)
            fmt.Println("adding", item)
            s.Add(item)
            wg.Done()
        }(i)
    }
    
    // Wait until all concurrent calls finished and print our set
    wg.Wait()
    fmt.Println(s)
}
