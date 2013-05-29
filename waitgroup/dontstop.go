package main

import "fmt"

// START OMIT

func doWorkInParallel(input int) {
	// Do some important work
	result := input + 1
	fmt.Printf("The result is %v\n", result)
}

func main() {
	for ii := 0; ii < 10; ii++ {
		go doWorkInParallel(ii)
	}
}

// END OMIT
