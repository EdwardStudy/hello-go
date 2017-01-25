package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// use default to non-blocking
	select {
	case msg := <-messages:
		fmt.Println("receivced message", msg)
	case sig := <-signals:
		fmt.Println("receivced signal", sig)
	default:
		fmt.Println("no activity")
	}
}
