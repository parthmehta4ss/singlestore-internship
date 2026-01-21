// Just some amazing behavior: try toggling the flag to find out the difference in behavior of code
package main

import (
	"fmt"
	"time"
)

func doWork(done chan int) {
	for {
		select {
		case <-done:
			fmt.Println("Received done signal, exiting goroutine.")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Still working...")
		}
	}
}
func main() {
	var flag bool = true // Toggle this flag to see different behaviors
	done := make(chan int)

	go doWork(done)
	time.Sleep(time.Second)
	close(done)
	if flag {
		time.Sleep(500 * time.Millisecond) // Give some time to see the goroutine exit
	}
}
