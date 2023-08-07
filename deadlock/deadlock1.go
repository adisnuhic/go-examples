package main

import "fmt"

func main() {
	c := make(chan int)

	// if we send a data from separate go routine it will avoid deadlock
	// go func() {
	// 	c <- 2
	// }()

	select {
	case <-c: // this case will never be selected because of deadlock - no one is sending the data
		fmt.Println("received") // this will never be printed
	}
}
