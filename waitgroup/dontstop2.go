package main

import (
	"fmt"
	"math/rand"
)

// START OMIT

func doWorkInParallel(input int) {
	// Do some important work
	coinflip := rand.Int() % 2
	if coinflip == 1 {
		doWorkInParallel(input - 1)
	}
	result := input + 1
	fmt.Printf("The result is %v\n", result)
}

func main() {
	for ii := 0; ii < 10; ii++ {
		go doWorkInParallel(ii)
	}
}

// END OMIT
