//http://www.capykoa.com/articles/10

package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "regexp"
    "strings"
)

var runStage = flag.String("stage", "", "specify the stage to run.  Can be 'mapper' or 'reducer'")

// Runs the mapper or reducer stage depending on the input
func main() {
    flag.Parse()
    if *runStage == "" {
        flag.PrintDefaults()
        return
    }

    switch *runStage {
    case "mapper":
        runMapper()
    case "reducer":
        runReducer()
    default:
        log.Fatalln("stage must be either 'mapper' or 'reducer'")
    }
}

// Reads from STDIN and write out each word in its own line. The word
// is lowercased and stripped of non-alphanumeric characters
func runMapper() {
    in := bufio.NewReader(os.Stdin)

    // Construct a regex to replace all non-alphanumeric characters
    reg, err := regexp.Compile("[^A-Za-z0-9]+")
    if err != nil {
        log.Fatal(err)
    }

    for {
        line, err := in.ReadString('\n')

        increment("wc_mapper", "lines")

        words := strings.Split(strings.TrimRight(line, "\n"), " ")

        for _, word := range words {
            strippedWord := reg.ReplaceAllString(word, "")

            fmt.Printf("%s\n", strings.ToLower(strippedWord))
        }

        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
    }
}

// The words are sorted when it is read from os.Stdin.
func runReducer() {
    in := bufio.NewReader(os.Stdin)

    var currentWord string
    var newWord string
    var count int

    // Loop until End Of File
    for {
        line, err := in.ReadString('\n')

        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        newWord = strings.TrimSuffix(line, "\n")

        // Because words from STDIN are sorted, we only need
        // to write to STDOUT whenever newWord != currentWord and reset the counter
        if currentWord == "" {
            currentWord = newWord
            count++
            continue
        } else if newWord != currentWord {
            // If we have a counter, we can easily keep track of internal metrics.
            // This counter is accessible in the jobtracker dashboard.

            increment("wc_reducer", "unique_words")
            fmt.Printf("%s\t%d\n", currentWord, count)

            count = 1
            currentWord = newWord
        } else {
            count++
        }
    }
}

// Per hadoop streaming docs, one can write to StdErr to have a report counter. Max of 120 counters, but 10-15 is recommended.
// http://hadoop.apache.org/docs/r0.18.3/streaming.html
func increment(group string, counter string) {
    fmt.Fprintf(os.Stderr, "reporter:counter:%s,%s,1\n", group, counter)
}
