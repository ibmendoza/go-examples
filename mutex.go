//https://medium.com/@deckarep/dancing-with-go-s-mutexes-92407ae927bf

package main

import (
	"fmt"
	"sync"
)

type DataStore struct {
	sync.Mutex // ? this mutex protects the cache below
	cache      map[string]string
}

func New() *DataStore {
	return &DataStore{
		cache: make(map[string]string),
	}
}

func (ds *DataStore) set(key string, value string) {
	ds.cache[key] = value
}

func (ds *DataStore) get(key string) string {
	if ds.count() > 0 {
		item := ds.cache[key]
		return item
	}
	return ""
}

func (ds *DataStore) count() int {
	return len(ds.cache)
}

func (ds *DataStore) Set(key string, value string) {
	ds.Lock()
	defer ds.Unlock()

	ds.set(key, value)
}

func (ds *DataStore) Get(key string) string {
	ds.Lock()
	defer ds.Unlock()

	return ds.get(key)
}

func (ds *DataStore) Count() int {
	ds.Lock()
	defer ds.Unlock()
	return ds.count()
}

func main() {
	store := New()
	store.Set("Go", "Lang")
	result := store.Get("Go")
	fmt.Println(result)
}
