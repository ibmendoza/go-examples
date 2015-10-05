//http://adampresley.com/2015/02/16/waiting-for-goroutines-to-finish-running-before-exiting.html

package main

import (
  "log"
  "os"
  "os/signal"
  "sync"
  "syscall"
)

func main() {
  log.Println("Starting application...")

  /*
* When SIGINT or SIGTERM is caught write to the quitChannel
*/
  quitChannel := make(chan os.Signal)
  signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

  shutdownChannel := make(chan bool)
  waitGroup := &sync.WaitGroup{}

  waitGroup.Add(1)

  /*
* Create a goroutine that does imaginary work
*/
  go func(shutdownChannel chan bool, waitGroup *sync.WaitGroup) {
    log.Println("Starting work goroutine...")
    defer waitGroup.Done()

    for {
      /*
* Listen on channels for message.
*/
      select {
      case _ = <-shutdownChannel:
        return

      default:
      }

      // Do some hard work here!
    }
  }(shutdownChannel, waitGroup)

  /*
* Wait until we get the quit message
*/
  <-quitChannel
  shutdownChannel <- true
  log.Println("Received quit. Sending shutdown and waiting on goroutines...")

  /*
* Block until wait group counter gets to zero
*/
  waitGroup.Wait()
  log.Println("Done.")
}
