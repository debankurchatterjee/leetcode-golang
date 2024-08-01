package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel for synchronization
	done := make(chan bool)

	go func() {
		for {
			fmt.Println("hi")
			time.Sleep(time.Second) // Pause for a second
			done <- true            // Signal that "hi" has been printed
			<-done                  // Wait for "hello" to be printed
		}
	}()

	go func() {
		for {
			<-done // Wait for "hi" to be printed
			fmt.Println("hello")
			time.Sleep(time.Second) // Pause for a second
			done <- true            // Signal that "hello" has been printed
		}
	}()

	// Prevent the main goroutine from exiting
	select {}
}
