//http://technosophos.com/2013/09/14/using-gos-built-logger-log-syslog.html
package main

import(
    "log"
    "log/syslog"
)

func main() {

    // Configure logger to write to the syslog. You could do this in init(), too.
    logwriter, e := syslog.New(syslog.LOG_NOTICE, "myprog")
    if e == nil {
        log.SetOutput(logwriter)
    }

    // Now from anywhere else in your program, you can use this:
    log.Print("Hello Logs!")
}
