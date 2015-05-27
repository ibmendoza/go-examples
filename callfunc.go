//Courtesy of https://github.com/golang-samples

package main

import (
	"errors"
	"fmt"
)

import "reflect"

func add(a, b int) int {
	return a + b
}

func sum(args ...int) int {
	s := 0
	for _, i := range args {
		s += i
	}
	return s
}

func divmod(a, b int) (div int, mod int) {
	div = a / b
	mod = a - div*b
	return
}

func CallAny(f interface{}, args ...interface{}) (out []interface{}, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("Failed to call")
		}
	}()
	fv := reflect.ValueOf(f)
	ft := fv.Type()
	margs := ft.NumIn()
	if ft.In(0).Name() == "" {
		margs = len(args)
	}
	inv := make([]reflect.Value, margs)
	for n := 0; n < margs; n++ {
		if n < len(args) {
			inv[n] = reflect.ValueOf(args[n])
		} else {
			inv[n] = reflect.Zero(ft.In(n))
		}
	}
	outv := fv.Call(inv)
	out = make([]interface{}, ft.NumOut())
	for n := 0; n < ft.NumOut(); n++ {
		out[n] = outv[n].Interface()
	}
	return out, nil
}

func main() {
	r1, e := CallAny(add, 1, 2) // Return 3
	if e == nil {
		fmt.Println(r1[0])
	}
	r2, e := CallAny(sum, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // Return 55
	if e == nil {
		fmt.Println(r2[0])
	}
	r3, e := CallAny(divmod, 7, 3) // Return 2, 1
	if e == nil {
		fmt.Println(r3[0], r3[1])
	}
}
