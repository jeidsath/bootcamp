package main

import "fmt"

var availableMemory chan int = make(chan int)
var results chan int = make(chan int)

func makeMemoryAvailable(amount int, limitchan chan bool) {
	for ii := 0; ii < cap(limitchan); ii++ {
		limitchan <- true
	}

	availableMemory <- amount
}

func trackMemory(amount int, availableMemory chan int) {
	mem := <-availableMemory
	mem += amount
	if mem < 0 {
		panic("Out of memory!")
	}
	availableMemory <- mem
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

var limitchan chan bool = make(chan bool, 3) // HL

func crunchNumber(number int, availableMemory, results chan int, limitchan chan bool) {
	<-limitchan // HL
	trackMemory(-100, availableMemory)
	results <- heavyProcessing(number)
	trackMemory(100, availableMemory)
	limitchan <- true // HL
}

func main() {
	numbers_to_crunch := [...]int{100, 303, 14, 12, 10010, 7}
	go makeMemoryAvailable(300, limitchan) // Prefills limitchan

	for _, number := range numbers_to_crunch {
		go crunchNumber(number, availableMemory, results, limitchan)
	}

	showResults(results, len(numbers_to_crunch))
}

// END OMIT
