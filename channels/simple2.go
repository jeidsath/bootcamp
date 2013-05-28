package main

import "fmt"

// START OMIT

func sayit(ss string, cc chan int) {
	for {
		<-cc // Receive from channel cc // HL
		fmt.Println(ss)
	}
}

func main() {
	cc := make(chan int)

	go sayit("Hey!", cc)
	go sayit("Ho!", cc)

	for ii := 0; ii < 10; ii++ {
		cc <- 1 // Send "1" to channel cc // HL
	}
}

// END OMIT
