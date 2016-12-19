//https://github.com/upitau/goinbigdata/blob/master/examples/calculator/calculator.go

package calculator

type Random interface {
	Random(limit int) int
}

type Calculator interface {
	Add(x, y int) int
	Subtract(x, y int) int
	Multiply(x, y int) int
	Divide(x, y int) int
	Random() int
}

func newCalculator(rnd Random) Calculator {
	return calc{
		rnd: rnd,
	}
}

type calc struct {
	rnd Random
}

func (c calc) Add(x, y int) int {
	return x + y
}

func (c calc) Subtract(x, y int) int {
	return x - y
}

func (c calc) Multiply(x, y int) int {
	return x * y
}

func (c calc) Divide(x, y int) int {
	return x / y
}

func (c calc) Random() int {
	return c.rnd.Random(100)
}
