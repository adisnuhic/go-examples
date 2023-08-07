package main

import "fmt"

func main() {
	fmt.Println("Channels are in Deadlock")

	ch1 := make(chan int)
	ch2 := make(chan int)

	// since there is no go routines that are sending data to channels main go routine remains blocked
	// As a result, you have a deadlock, and the program won't proceed beyond this point.
	<-ch1
	<-ch2
}
