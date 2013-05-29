package main

import "fmt"

var mem chan int = make(chan int)
var results chan int = make(chan int)

func makeMemoryAvailable(amount int) {
	mem <- amount
}

func trackMemory(amount int, mem chan int) {
	memory := <-mem
	memory += amount
	if memory < 0 {
		panic("Out of memory!")
	}
	mem <- memory
}

func showResults(results chan int, length int) {
	for ii := 0; ii < length; ii++ {
		rr := <-results
		fmt.Printf("%d\n", rr)
	}
}

func heavyProcessing(number int) int {
	return number + 1
}

// START OMIT

func crunchNumber(number int, mem, results chan int) {
	trackMemory(-100, mem)
	results <- heavyProcessing(number)
	trackMemory(100, mem)
}

func main() {
	numbers_to_crunch := [...]int{100, 303, 14, 12, 10010, 7}
	go makeMemoryAvailable(300) // Not enough memory to run it all at once!

	for _, number := range numbers_to_crunch {
		go crunchNumber(number, mem, results)
	}

	showResults(results, len(numbers_to_crunch))
}

// END OMIT
