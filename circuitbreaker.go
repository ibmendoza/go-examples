package main

import (
	"errors"
	"fmt"
	"github.com/peterbourgon/g2s"
	cb "github.com/rubyist/circuitbreaker"
	"io/ioutil"
	"log"
	"time"
)

func ExampleErrorThresholdBreaker() {
	// This example sets up a ThresholdBreaker that will trip if remoteCall returns
	// an error 10 times in a row. The error returned by Call() will be the error
	// returned by remoteCall, unless the breaker has been tripped, in which case
	// it will return ErrBreakerOpen.
	breaker := cb.NewThresholdBreaker(10)
	err := breaker.Call(errorCall, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleThresholdBreaker() {
	// This example sets up a ThresholdBreaker that will trip if remoteCall returns
	// an error 10 times in a row. The error returned by Call() will be the error
	// returned by remoteCall, unless the breaker has been tripped, in which case
	// it will return ErrBreakerOpen.
	breaker := cb.NewThresholdBreaker(10)
	err := breaker.Call(remoteCall, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleTimeoutBreaker() {
	// This example sets up a TimeoutBreaker that will trip if remoteCall returns
	// an error OR takes longer than one second 10 times in a row. The error returned
	// by Call() will be the error returned by remoteCall with two exceptions: if
	// remoteCall takes longer than one second the return value will be ErrBreakerTimeout,
	// if the breaker has been tripped the return value will be ErrBreakerOpen.
	breaker := cb.NewThresholdBreaker(10)
	err := breaker.Call(remoteCall, time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleConsecutiveBreaker() {
	// This example sets up a FrequencyBreaker that will trip if remoteCall returns
	// an error 10 times in a row within a period of 2 minutes.
	breaker := cb.NewConsecutiveBreaker(10)
	err := breaker.Call(remoteCall, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleHTTPClient() {
	// This example sets up an HTTP client wrapped in a TimeoutBreaker. The breaker
	// will trip with the same behavior as TimeoutBreaker.
	client := cb.NewHTTPClient(time.Second*5, 10, nil)

	resp, err := client.Get("http://example.com/resource.json")
	if err != nil {
		log.Fatal(err)
	}
	resource, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", resource)
}

func ExampleBreaker_events() {
	// This example demonstrates the BreakerTripped and BreakerReset callbacks. These are
	// available on all breaker types.
	breaker := cb.NewThresholdBreaker(10)
	events := breaker.Subscribe()

	go func() {
		for {
			e := <-events
			switch e {
			case cb.BreakerTripped:
				log.Println("breaker tripped")
			case cb.BreakerReset:
				log.Println("breaker reset")
			case cb.BreakerFail:
				log.Println("breaker fail")
			case cb.BreakerReady:
				log.Println("breaker ready")
			}
		}
	}()

	breaker.Fail()
	//breaker.Reset()
}

func ExamplePanel() {
	// This example demonstrates using a Panel to aggregate and manage circuit breakers.
	breaker1 := cb.NewThresholdBreaker(10)
	breaker2 := cb.NewRateBreaker(0.95, 100)

	panel := cb.NewPanel()
	panel.Add("breaker1", breaker1)
	panel.Add("breaker2", breaker2)

	// Elsewhere in the code ...
	b1, _ := panel.Get("breaker1")
	b1.Call(func() error {
		// Do some work
		return nil
	}, 0)

	b2, _ := panel.Get("breaker2")
	b2.Call(func() error {
		// Do some work
		return nil
	}, 0)
}

func ExamplePanel_stats() {
	// This example demonstrates how to push circuit breaker stats to statsd via a Panel.
	// This example uses g2s. Anything conforming to the Statter interface can be used.
	s, err := g2s.Dial("udp", "statsd-server:8125")
	if err != nil {
		log.Fatal(err)
	}

	breaker := cb.NewThresholdBreaker(10)
	panel := cb.NewPanel()
	panel.Statter = s
	panel.StatsPrefixf = "sys.production.%s"
	panel.Add("x", breaker)

	breaker.Trip()  // sys.production.circuit.x.tripped
	breaker.Reset() // sys.production.circuit.x.reset, sys.production.circuit.x.trip-time
	breaker.Fail()  // sys.production.circuit.x.fail
	breaker.Ready() // sys.production.circuit.x.ready (if it's tripped and ready to retry)
}

func remoteCall() error {
	// Expensive remote call
	return nil
}

func errorCall() error {
	// Expensive remote call
	return errors.New("error")
}

func main() {
	breaker := cb.NewThresholdBreaker(10)

	i := 1
	for i <= 12 {
		if breaker.Ready() {
			log.Println(breaker.Call(errorCall, 0))
		} else {
			log.Println("circuit breaker is open")
		}
		i++
	}
}
