package main

import (
	"fmt"
	"github.com/joyrexus/buckets"
)

func main() {
	bx, _ := buckets.Open("bolt")
	//defer os.Remove(bx.Path())
	defer bx.Close()

	// Delete any existing bucket named "years".
	bx.Delete([]byte("years"))

	// Create a new bucket named "years".
	years, _ := bx.New([]byte("years"))

	// Setup items to insert in `years` bucket
	items := []struct {
		Key, Value []byte
	}{
		{[]byte("1970"), []byte("70")},
		{[]byte("1975"), []byte("75")},
		{[]byte("1980"), []byte("80")},
		{[]byte("1985"), []byte("85")},
		{[]byte("1990"), []byte("90")}, // min = 1990
		{[]byte("1995"), []byte("95")}, // min < 1995 < max
		{[]byte("2000"), []byte("00")}, // max = 2000
		{[]byte("2005"), []byte("05")},
		{[]byte("2010"), []byte("10")},
	}

	// Insert 'em.
	if err := years.Insert(items); err != nil {
		fmt.Printf("could not insert items in `years` bucket: %v\n", err)
	}

	// Time range to map over: 1990 <= key <= 2000.
	min := []byte("1990")
	max := []byte("2000")

	// Setup slice of items to collect results.
	type item struct {
		Key, Value []byte
	}
	results := []item{}

	// Anon func to map over matched keys.
	do := func(k, v []byte) error {
		results = append(results, item{k, v})
		return nil
	}

	if err := years.MapRange(do, min, max); err != nil {
		fmt.Printf("could not map items within range: %v\n", err)
	}

	for _, item := range results {
		fmt.Printf("%s -> %s\n", item.Key, item.Value)
	}
}
