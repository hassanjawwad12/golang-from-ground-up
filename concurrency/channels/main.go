package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	done <- true
	close(done) // channel cant be used again
}

func main() {
	// dones := make([]chan bool, 4)

	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)

	// go greet("Nice to meet you!", dones[0])
	// go greet("How are you?", dones[1])
	// go slowGreet("How ... are ... you ...?", dones[2])
	// go greet("I hope you're liking the course!", dones[3])

	// for _, done := range dones {
	// 	<-done
	// }

	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	for range done {
	}

}
