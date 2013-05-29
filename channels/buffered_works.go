package main

import "fmt"

var mem chan int = make(chan int)
var results chan int = make(chan int)

func makeMemoryAvailable(amount int, lim chan bool) {
	for ii := 0; ii < cap(lim); ii++ {
		lim <- true
	}

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

var lim chan bool = make(chan bool, 3) // HL

func crunchNumber(number int, mem, results chan int, lim chan bool) {
	<-lim // HL
	trackMemory(-100, mem)
	results <- heavyProcessing(number)
	trackMemory(100, mem)
	lim <- true // HL
}

func main() {
	numbers_to_crunch := [...]int{100, 303, 14, 12, 10010, 7}
	go makeMemoryAvailable(300, lim) // Pre-fills lim

	for _, number := range numbers_to_crunch {
		go crunchNumber(number, mem, results, lim)
	}

	showResults(results, len(numbers_to_crunch))
}

// END OMIT
