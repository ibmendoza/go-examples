//http://marcio.io/2015/07/cheap-mapreduce-in-go/

package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	MaxWorkers = 10
)

type Telemetry struct {
	Request struct {
    	Sender  string `json:"Sender,omitempty"`
    	Trigger string `json:"Trigger,omitempty"`
    } `json:"Request,omitempty"`

	App struct {
    	Program  string `json:"Program,omitempty"`
    	Build    string `json:"Build,omitempty"`
    	License  string `json:"License,omitempty"`
    	Version  string `json:"Version,omitempty"`
    } `json:"App,omitempty"`

	Connection struct {
    	Type string `json:"Type,omitempty"`
    } `json:"Connection,omitempty"`

	Region struct {
    	Continent string `json:"Continent,omitempty"`
    	Country   string `json:"Country,omitempty"`
    } `json:"Region,omitempty"`

	Client struct {
    	OsVersion    string `json:"OsVersion,omitempty"`
    	Language     string `json:"Language,omitempty"`
    	Architecture string `json:"Architecture,omitempty"`
    } `json:"Client,omitempty"`
}

func enumerateFiles(dirname string) chan interface{} {
	output := make(chan interface{})
	go func() {
		filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				output <- path
			}
			return nil
		})
		close(output)
	}()
	return output
}

func enumerateJSON(filename string) chan string {
	output := make(chan string)
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}

			// ignore any meta comments on top of JSON file
			if strings.HasPrefix(line, "#") == true {
				continue
			}

			// add each json line to our enumeration channel
			output <- line
		}
		close(output)
	}()
	return output
}

// MapperCollector is a channel that collects the output from mapper tasks
type MapperCollector chan chan interface{}

// MapperFunc is a function that performs the mapping part of the MapReduce job
type MapperFunc func(interface{}, chan interface{})

// ReducerFunc is a function that performs the reduce part of the MapReduce job
type ReducerFunc func(chan interface{}, chan interface{})

func mapperDispatcher(mapper MapperFunc, input chan interface{}, collector MapperCollector) {
	for item := range input {
		taskOutput := make(chan interface{})
		go mapper(item, taskOutput)
		collector <- taskOutput
	}
	close(collector)
}

func reducerDispatcher(collector MapperCollector, reducerInput chan interface{}) {
	for output := range collector {
		reducerInput <- <-output
	}
	close(reducerInput)
}

func mapper(filename interface{}, output chan interface{}) {
	results := map[Telemetry]int{}

    // start the enumeration of each JSON lines in the file
	for line := range enumerateJSON(filename.(string)) {

		// decode the telemetry JSON line
		dec := json.NewDecoder(strings.NewReader(line))
		var telemetry Telemetry

		// if line cannot be JSON decoded then skip to next one
		if err := dec.Decode(&telemetry); err == io.EOF {
			continue
		} else if err != nil {
			continue
		}

		// stores Telemetry structure in the mapper results dictionary
		previousCount, exists := results[telemetry]
		if !exists {
			results[telemetry] = 1
		} else {
			results[telemetry] = previousCount + 1
		}
	}

	output <- results
}

func reducer(input chan interface{}, output chan interface{}) {
	results := map[Telemetry]int{}
	for matches := range input {
		for key, value := range matches.(map[Telemetry]int) {
			_, exists := results[key]
			if !exists {
				results[key] = value
			} else {
				results[key] = results[key] + value
			}
		}
	}
	output <- results
}

func mapReduce(mapper MapperFunc, reducer ReducerFunc, input chan interface{}) interface{} {

	reducerInput := make(chan interface{})
	reducerOutput := make(chan interface{})
	mapperCollector := make(MapperCollector, MaxWorkers)

	go reducer(reducerInput, reducerOutput)
	go reducerDispatcher(mapperCollector, reducerInput)
	go mapperDispatcher(mapper, input, mapperCollector)

	return <-reducerOutput
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Processing. Please wait....")

    // start the enumeration of files to be processed into a channel
	input := enumerateFiles(".")

    // this will start the map reduce work
	results := mapReduce(mapper, reducer, input)

	// open output file
	f, err := os.Create("telemetry.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// make a write buffer
	writer := csv.NewWriter(f)

	for telemetry, value := range results.(map[Telemetry]int) {

		var record []string

		record = append(record, telemetry.Request.Sender)
		record = append(record, telemetry.Request.Trigger)
		record = append(record, telemetry.App.Program)
		record = append(record, telemetry.App.Build)
		record = append(record, telemetry.App.License)
		record = append(record, telemetry.App.Version)
		record = append(record, telemetry.Connection.Type)
		record = append(record, telemetry.Region.Continent)
		record = append(record, telemetry.Region.Country)
		record = append(record, telemetry.Client.OsVersion)
		record = append(record, telemetry.Client.Language)
		record = append(record, telemetry.Client.Architecture)

        	// The last field of the CSV line is the aggregate count for each occurrence
		record = append(record, strconv.Itoa(value))

		writer.Write(record)
	}

	writer.Flush()

	fmt.Println("Done!")
}
