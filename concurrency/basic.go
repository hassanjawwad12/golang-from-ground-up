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
	done <- true // sends data (true) to the channel
}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done) // this waits for the slowgreet to finish execution
	<-done
	<-done
	<-done
	<-done
	// read from the channel , we are waiting for the channle to omit data
	// go will only continue after data in done is received
}
