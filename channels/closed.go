package main

import "fmt"

// START OMIT

func main() {
	cc := make(chan int)

	close(cc)

	for ii := 0; ii < 5; ii++ {
		value, err := <-cc
		fmt.Printf("%v, %v\n", value, err)
	}
}

// END OMIT
