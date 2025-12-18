/*
Implement the dining philosopher’s problem with the following constraints:

- There are 5 philosophers sitting around a table with 5 shared chopsticks,
  one chopstick between each adjacent pair of philosophers.

- Each philosopher eats exactly 3 times.

- Philosophers are allowed to pick up chopsticks in either order
  (not constrained to lowest-numbered first).

- A "host" goroutine controls access to the table and allows
  no more than 2 philosophers to eat concurrently.

- Each philosopher is numbered from 1 through 5.

- When a philosopher starts eating (after acquiring both chopsticks),
  it prints:  "starting to eat <number>"

- When a philosopher finishes eating (before releasing the chopsticks),
  it prints:  "finishing eating <number>"
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	id                    int
	leftChopstick         *chopstick
	rightChopstick        *chopstick
	hostRequestChan       chan chan struct{}
	hostFinishedEatingChan chan struct{}
}

// dine represents the life of a philosopher: think → request permission → eat (3 times).
func (p philosopher) dine(eatTimes int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < eatTimes; i++ {
		// Request permission from the host to start eating.
		permission := make(chan struct{})
		p.hostRequestChan <- permission
		<-permission // Block until the host grants permission.

		// Pick up chopsticks in a non-uniform order.
		// For example, even-numbered philosophers pick up right then left,
		// odd-numbered philosophers pick up left then right.
		first, second := p.leftChopstick, p.rightChopstick
		if p.id%2 == 0 {
			first, second = p.rightChopstick, p.leftChopstick
		}

		first.Lock()
		second.Lock()

		// At this point, the philosopher has acquired both chopsticks.
		fmt.Printf("starting to eat %d\n", p.id)
		// Optional: simulate eating time to make interleavings more visible.
		time.Sleep(10 * time.Millisecond)

		// Print finishing message before releasing locks, as required.
		fmt.Printf("finishing eating %d\n", p.id)

		second.Unlock()
		first.Unlock()

		// Inform the host that this philosopher has finished eating this round.
		p.hostFinishedEatingChan <- struct{}{}
	}
}

// host runs in its own goroutine and enforces the "at most 2 philosophers eating" rule.
func host(requests <-chan chan struct{}, finished <-chan struct{}, totalMeals int) {
	currentlyEating := 0
	completedMeals := 0

	for completedMeals < totalMeals {
		// Only accept new eating requests when fewer than 2 philosophers are eating.
		var reqChan <-chan chan struct{}
		if currentlyEating < 2 {
			reqChan = requests
		} else {
			// Disable the request case when the table is "full".
			reqChan = nil
		}

		select {
		// Grant permission to a philosopher if there is capacity.
		case perm := <-reqChan:
			currentlyEating++
			perm <- struct{}{} // Grant permission.

		// A philosopher has finished eating one meal.
		case <-finished:
			currentlyEating--
			completedMeals++
		}
	}
}

func main() {
	const philosopherCount = 5
	const eatTimes = 3
	totalMeals := philosopherCount * eatTimes

	// Shared channels for communication with the host.
	hostRequestChan := make(chan chan struct{})
	hostFinishedEatingChan := make(chan struct{})

	// Start the host goroutine.
	go host(hostRequestChan, hostFinishedEatingChan, totalMeals)

	// Create chopsticks.
	chopsticks := make([]*chopstick, philosopherCount)
	for i := 0; i < philosopherCount; i++ {
		chopsticks[i] = &chopstick{}
	}

	// Create philosophers and assign chopsticks.
	var wg sync.WaitGroup
	philosophers := make([]*philosopher, philosopherCount)
	for i := 0; i < philosopherCount; i++ {
		philosophers[i] = &philosopher{
			id:                    i + 1, // philosophers are numbered 1..5
			leftChopstick:         chopsticks[i],
			rightChopstick:        chopsticks[(i+1)%philosopherCount],
			hostRequestChan:       hostRequestChan,
			hostFinishedEatingChan: hostFinishedEatingChan,
		}
	}

	// Launch each philosopher in its own goroutine.
	for _, p := range philosophers {
		wg.Add(1)
		go p.dine(eatTimes, &wg)
	}

	// Wait until all philosophers have completed their 3 meals.
	wg.Wait()
