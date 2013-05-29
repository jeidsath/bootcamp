package main

import "fmt"

// START OMIT

func doWorkInParallel(input int, cc chan bool) {
	// Do some important work
	result := input + 1
	fmt.Printf("The result is %v\n", result)
	cc <- true
}

func main() {
	const processes int = 10
	cc := make(chan bool, processes)
	for ii := 0; ii < processes; ii++ {
		go doWorkInParallel(ii, cc)
	}
	for ii := 0; ii < processes; ii++ {
		<-cc
	}
}

// END OMIT
