package main

import (
	"fmt"
	"sync"
)

var counter int

func increment() {

	counter++
	fmt.Printf("Counter: %d\n", counter)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("I am from goroute %d\n", i)
			increment()
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("All works completed")
}
