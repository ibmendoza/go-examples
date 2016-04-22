//https://github.com/jehiah/gomrjob/blob/master/example/example_mr.go

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jehiah/gomrjob"
	"github.com/jehiah/gomrjob/hdfs"
	"github.com/jehiah/lru"
)

var (
	input = flag.String("input", "", "path to hdfs input file")
)

type JsonEntryCounter struct {
	KeyField string
}

// An example Map function. It consumes json data and yields a value for each line
func (s *JsonEntryCounter) Mapper(r io.Reader, w io.Writer) error {
	log.Printf("map_input_file %s", os.Getenv("map_input_file"))
	wg, out := gomrjob.JsonInternalOutputProtocol(w)

	// for efficient counting, use an in-memory counter that flushes the least recently used item
	// less Mapper output makes for faster sorting and reducing.
	counter := lru.NewLRUCounter(func(k interface{}, v int64) {
		out <- gomrjob.KeyValue{k, v}
	}, 1)

	for data := range gomrjob.JsonInputProtocol(r) {
		gomrjob.Counter("example_mr", "Map Lines Read", 1)
		key, err := data.Get(s.KeyField).String()
		if err != nil {
			gomrjob.Counter("example_mr", "Missing Key", 1)
		} else {
			counter.Incr(key, 1)
		}
	}
	counter.Flush()
	close(out)
	wg.Wait()
	return nil
}

// just re-use the reducer as the combiner
func (s *JsonEntryCounter) Combiner(r io.Reader, w io.Writer) error {
	return s.Reducer(r, w)
}

// // A simple reduce function that counts keys
// func (s *MRStep) Reducer(r io.Reader, w io.Writer) error {
// 	wg, out := gomrjob.JsonInternalOutputProtocol(w)
// 	for kv := range gomrjob.JsonInternalInputProtocol(r) {
// 		var i int64
// 		for v := range kv.Values {
// 			vv, err := v.Int64()
// 			if err != nil {
// 				gomrjob.Counter("example_mr", "non-int value", 1)
// 				log.Printf("non-int value %s", err)
// 			} else {
// 				i += vv
// 			}
// 		}
// 		keyString, err := kv.Key.String()
// 		if err != nil {
// 			gomrjob.Counter("example_mr", "non-string key", 1)
// 			log.Printf("non-string key %s", err)
// 		}
// 		out <- gomrjob.KeyValue{keyString, i}
// 	}
// 	close(out)
// 	wg.Wait()
// 	return nil
// }

func (s *JsonEntryCounter) Reducer(r io.Reader, w io.Writer) error {
	wg, out := gomrjob.RawJsonInternalOutputProtocol(w)
	for kv := range gomrjob.RawJsonInternalInputProtocol(r) {
		var i int64
		for v := range kv.Values {
			vv, err := v.Int64()
			if err != nil {
				gomrjob.Counter("example_mr", "non-int value", 1)
				log.Printf("non-int value %s", err)
			} else {
				i += vv
			}
		}
		out <- gomrjob.KeyValue{kv.Key, i}
	}
	close(out)
	wg.Wait()
	return nil
}

func main() {
	flag.Parse()

	runner := gomrjob.NewRunner()
	runner.Name = "test-gomrjob"
	runner.InputFiles = append(runner.InputFiles, *input)
	runner.ReducerTasks = 3
	runner.Steps = append(runner.Steps, &JsonEntryCounter{"api_path"})
	err := runner.Run()
	if err != nil {
		gomrjob.Status(fmt.Sprintf("Run error %s", err))
		log.Fatalf("Run error %s", err)
	}
	cmd := hdfs.Cat(fmt.Sprintf("%s/part-*", runner.Output))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run() // err?
	// runner.Cleanup()

}
