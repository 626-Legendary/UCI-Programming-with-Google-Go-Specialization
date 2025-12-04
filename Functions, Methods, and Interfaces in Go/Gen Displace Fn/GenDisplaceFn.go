package main

import (
	"fmt"
)

// GenDisplaceFn returns a function that computes displacement as a function of time t.
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so float64
	var t float64

	fmt.Print("Enter acceleration a: ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity vo: ")
	fmt.Scan(&vo)

	fmt.Print("Enter initial displacement so: ")
	fmt.Scan(&so)

	// Generate displacement function
	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Enter time t: ")
	fmt.Scan(&t)

	// Compute displacement
	fmt.Println("Displacement =", fn(t))
}
