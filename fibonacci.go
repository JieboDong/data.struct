package main

import (
	"fmt"
)

func Fibonacci() func() int {

	x, y := 0, 1

	return func() int {
		next := x + y
		ff := x
		x = y
		y = next
		return ff
	}

}

const DEPATH = 50

func main() {

	var ouput [DEPATH]int
	f := Fibonacci()
	for i := 0; i < DEPATH; i++ {

		ouput[i] = f()

	}
	fmt.Println(ouput)
}
