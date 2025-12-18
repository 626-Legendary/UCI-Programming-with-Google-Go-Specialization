/*
Write two goroutines that exhibit a data race when executed concurrently.
The race occurs because both goroutines access and modify the same shared
variable `x` without any synchronization, leading to non-deterministic results.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var x int

	// WaitGroup is used to wait for all launched goroutines to complete.
	var w sync.WaitGroup

	// Launch two goroutines that concurrently increment the shared variable x.
	for i := 0; i < 2; i++ {
		// Increment the WaitGroup counter to track one additional goroutine.
		w.Add(1)

		// Start an anonymous goroutine.
		go func() {
			defer w.Done() // Ensure the WaitGroup counter is decremented when the goroutine exits.

			for j := 0; j < 10; j++ {
				/*
					x++ is a non-atomic read-modify-write operation.

					Each increment conceptually does:
						1. Read the current value of x.
						2. Add 1 to the value.
						3. Write the new value back to x.

					When two goroutines execute x++ concurrently without synchronization,
					their read and write operations can be interleaved in arbitrary ways.
					This causes a data race and may result in some increments being
					logically "lost," so the final value of x is not guaranteed.
				*/
				x++
				time.Sleep(1 * time.Millisecond) // Artificial delay to increase the likelihood of interleaving.
			}
		}()
	}

	// Block until both goroutines have finished executing.
	w.Wait()

	/*
		Expected value without a race:
			2 goroutines × 10 increments each = 20

		Due to the data race, the actual result may be less than or equal to 20
		and is non-deterministic across different runs.
	*/
	fmt.Println("Result:", x)
}
