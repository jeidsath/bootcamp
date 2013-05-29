package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// START OMIT

func doWorkInParallel(input int, wg *sync.WaitGroup) {
	defer wg.Done() // HL
	// Do some important work
	coinflip := rand.Int() % 2
	if coinflip == 1 {
		wg.Add(1) // HL
		doWorkInParallel(input - 1, wg)
	}
	result := input + 1
	fmt.Printf("The result is %v\n", result)
}

func main() {
	wg := sync.WaitGroup{} // HL
	for ii := 0; ii < 10; ii++ {
		wg.Add(1) // HL
		go doWorkInParallel(ii, &wg)
	}
	wg.Wait() // HL
}

// END OMIT
