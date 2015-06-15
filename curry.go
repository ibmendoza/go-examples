//http://play.golang.org/p/6eVRXjYozz

//http://stackoverflow.com/questions/19394868/how-can-go-lang-curry

package main

import (
    "fmt"
)

func mkAdd(a int) func(...int) int {
    return func(b... int) int {
        for _, i := range b {
            a += i
        }
        return a
    }
}

func main() {
    add2 := mkAdd(2)
    add3 := mkAdd(3)
    fmt.Println(add2(5,3), add3(6))
}
