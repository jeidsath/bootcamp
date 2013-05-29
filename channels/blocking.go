package main

import (
	"fmt"
)

// START OMIT

func bar(cc, qq chan int) {
	<-cc // Block waiting for func foo // HL
	fmt.Print("bar\n")
	qq <- 1
}

func foo(cc chan int) {
	fmt.Print("foo")
	cc <- 1
}

func main() {
	cc := make(chan int)
	qq := make(chan int)
	go bar(cc, qq) // It doesn't matter that we start bar or foo first // HL
	go foo(cc)     // Output will be "foobar" // HL

	<-qq // Block waiting for func bar // HL
}

// END OMIT
