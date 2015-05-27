//https://itjumpstart.wordpress.com/2014/11/21/channels-versus-closure-variables

package main
import (
    "fmt"
    "strconv"
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
//a function that returns a channel (like a promise object)
func fibo1(arg int) <-chan int {
    c := make(chan int)
    go func(i int) {
        c <- fibo(i)
    }(arg)
    return c
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
    sum = <-fibo1(21) + <-fibo1(14) + <-fibo1(7)
    fmt.Println("Sum of three fibo functions: " + strconv.Itoa(sum))
}
