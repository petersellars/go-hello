// This sample program demonstrates how to create goroutines and
// how the scheduler behaves.
package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

// main is the entrypoint for all Go programs
func main() {

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines
	fmt.Println("Start Goroutines")
	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// printPrime displays prime number for the first 5000 numbers
func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done
	defer wg.Done()

	next:
		for outer := 2; outer < 5000; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer%inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}
	fmt.Println("Completed", prefix)
}