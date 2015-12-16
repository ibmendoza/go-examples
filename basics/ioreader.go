//https://medium.com/@mschuett/golangs-reader-interface-bd2917d5ce83#
//http://play.golang.org/p/ejpUVOx8jR

package main

import "io"
import "io/ioutil"
import "log"

type Reader struct {
	read string
	done bool
}

func NewReader(toRead string) *Reader {
	return &Reader{toRead, false}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.done {
		return 0, io.EOF
	}
	for i, b := range []byte(r.read) {
		p[i] = b
	}
	r.done = true
	return len(r.read), nil
}

func main() {
	M := NewReader("test")
	stuff, _ := ioutil.ReadAll(M)
	log.Printf("%s", stuff)
}
