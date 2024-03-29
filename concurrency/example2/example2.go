// Sample program to show how to use the WithCancel function.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context that is cancellable only manually.
	// The cancel function must be called regardless of the outcome.
	ctx, cancel := context.WithCancel(context.Background())

	// Ask the goroutine to do some work for us.
	go func() {
		// Wait for the work to finish. If it takes too long move on.
		select {
		case <-time.After(100 * time.Millisecond):
			fmt.Println("moving on")
		case <-ctx.Done():
			fmt.Println("work complete")
		}
	}()

	// Simulate work.
	time.Sleep(50 * time.Millisecond)

	// Report the work is done.
	cancel()

	// Just hold the program to see the output.
	time.Sleep(time.Second)
}
