package main

func main() {
	// START OMIT
	// There is no place to put anything in cc1!
	cc1 := make(chan int, 0) // Unbuffered channel, capacity 0

	// Lots of room in cc2, but why would I want to 
	// store anything in a channel?
	cc2 := make(chan int, 10) // Buffered channel, capacity 10
	// END OMIT
}
