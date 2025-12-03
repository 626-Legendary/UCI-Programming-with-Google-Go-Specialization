package main

import (
	"fmt"
)

func main() {
	fmt.Scan()
	// s = 1/2at^2 +Vt+s0
	GenDisplaceFn()
}

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {

}
