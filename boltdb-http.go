http://stackoverflow.com/questions/30707208/boltdb-key-value-data-store-purely-in-go
package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"

    "github.com/boltdb/bolt"
    "github.com/gorilla/mux"
)

type server struct {
    db *bolt.DB
}

func newServer(filename string) (s *server, err error) {
    s = &server{}
    s.db, err = bolt.Open(filename, 0600, &bolt.Options{Timeout: 1 * time.Second})
    return
}

func (s *server) Put(bucket, key, contentType string, val []byte) error {
    return s.db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists([]byte(bucket))
        if err != nil {
            return err
        }
        if err = b.Put([]byte(key), val); err != nil {
            return err
        }
        return b.Put([]byte(fmt.Sprintf("%s-ContentType", key)), []byte(contentType))
    })
}

func (s *server) Get(bucket, key string) (ct string, data []byte, err error) {
    s.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        r := b.Get([]byte(key))
        if r != nil {
            data = make([]byte, len(r))
            copy(data, r)
        }

        r = b.Get([]byte(fmt.Sprintf("%s-ContentType", key)))
        ct = string(r)
        return nil
    })
    return
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    if vars["bucket"] == "" || vars["key"] == "" {
        http.Error(w, "Missing bucket or key", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case "POST", "PUT":
        data, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        err = s.Put(vars["bucket"], vars["key"], r.Header.Get("Content-Type"), data)
        w.WriteHeader(http.StatusOK)
    case "GET":
        ct, data, err := s.Get(vars["bucket"], vars["key"])
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Add("Content-Type", ct)
        w.Write(data)
    }
}

func main() {
    var (
        addr   string
        dbfile string
    )

    flag.StringVar(&addr, "l", ":9988", "Address to listen on")
    flag.StringVar(&dbfile, "db", "/var/data/bolt.db", "Bolt DB file")
    flag.Parse()

    log.Println("Using Bolt DB file:", dbfile)
    log.Println("Listening on:", addr)

    server, err := newServer(dbfile)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }

    router := mux.NewRouter()
    router.Handle("/v1/buckets/{bucket}/keys/{key}", server)
    http.Handle("/", router)

    log.Fatal(http.ListenAndServe(addr, nil))
}
