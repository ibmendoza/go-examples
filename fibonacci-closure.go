//https://itjumpstart.wordpress.com/2014/11/21/channels-versus-closure-variables

package main
import (
    "fmt"
    "strconv"
    "sync"
)
func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}
func fibo(n int) (result int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        result = x
        x, y = y, x+y
    }
    //close(c)
    return result
}
// fib returns a function that returns
// successive Fibonacci numbers.
//http://golang.org/doc/play/fib.go
func fib() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}
func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    sum := 0
    for i := range c {
        sum += i
        fmt.Println(i)
    }
    fmt.Println("Sum=" + strconv.Itoa(sum))
    //fmt.Println(fibo(10))
    var wg sync.WaitGroup
    var f1, f2, f3 int
    wg.Add(3)
    go func(arg int) {
        defer wg.Done()
        f1 = fibo(arg)
        fmt.Println("Goroutine1: " + strconv.Itoa(f1))
    }(21)
    go func(arg int) {
        defer wg.Done()
        f2 = fibo(arg)
        fmt.Println("Goroutine2: " + strconv.Itoa(f2))
    }(14)
    go func(arg int) {
        defer wg.Done()
        f3 = fibo(arg)
        fmt.Println("Goroutine3: " + strconv.Itoa(f3))
    }(7)
    wg.Wait()
    sum = f1 + f2 + f3
    fmt.Println("Sum of three fibo functions: " + strconv.Itoa(sum))
    fmt.Scanln()
}
