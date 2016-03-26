// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // Time parsing uses the same layout values as `Format`.
    t1, _ := time.Parse(
        time.RFC3339,
        "2016-03-26T14:07:32.7335156+08:00")
    p(t1)
    p(t1.Unix())
}
