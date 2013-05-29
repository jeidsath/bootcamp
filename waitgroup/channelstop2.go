package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT

func doWorkInParallel(input int, cc chan int) {
	// Do some important work
	coinflip := rand.Int() % 2
	if coinflip == 1 {
		launchedProcesses := <-cc
		launchedProcesses++
		cc <- launchedProcesses
		doWorkInParallel(input-1, cc)
	}
	result := input + 1
	fmt.Printf("The result is %v\n", result)
	launchedProcesses := <-cc
	launchedProcesses--
	cc <- launchedProcesses
}

// END OMIT

// START2 OMIT

func main() {
	cc := make(chan int, 1)
	cc <- 0
	for ii := 0; ii < 10; ii++ {
		launchedProcesses := <-cc
		launchedProcesses++
		cc <- launchedProcesses
		go doWorkInParallel(ii, cc)
	}

	for {
		select {
		case launchedProcesses := <-cc:
			if launchedProcesses == 0 {
				fmt.Println("All done!")
				return
			}
			cc <- launchedProcesses
			time.Sleep(time.Second)
		}
	}
}

// END2 OMIT
