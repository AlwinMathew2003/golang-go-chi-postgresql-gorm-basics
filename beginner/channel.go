package main

import (
	"fmt"
)

func producer(id int, receiver chan<- string) {
	receiver <- fmt.Sprintf("Produced send a message %d", id)
}

func main() {
	ch := make(chan string, 5)

	for i := 1; i <= 5; i++ {
		go producer(i, ch)
		fmt.Printf("Hai %d\n", i)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Received")
		fmt.Println(<-ch)
	}
}
