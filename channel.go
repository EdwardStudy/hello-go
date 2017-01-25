package main

import (
	"fmt"
)

import "time"

func worker(done chan bool) {
	// messages := make(chan string, 2)
	// messages <- "buffered"
	// messages <- "channel"

	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

// 使用 `make(chan val-type)` 创建一个新的通道。

//send // 使用 `channel <-` 语法 _发送_ 一个新的值到通道中。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//recv // 使用 `channel <-` 语法 _发送_ 一个新的值到通道中。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// // 使用 `make(chan val-type)` 创建一个新的通道。
	// messages := make(chan string, 2)
	// messages <- "buffered"
	// messages <- "channel"

	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// done := make(chan bool, 1)
	// go worker(done)

	// //block
	// <-done
	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)
	// ping(pings, "passed message")
	// pong(pings, pongs)
	// fmt.Println(<-pongs)

	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在若干时间后接收一个值，这个用来模拟例如
	// 并行的 Go 协程中阻塞的 RPC 操作
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// 我们使用 `select` 关键字来同时等待这两个值，并打
	// 印各自接收到的值。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 使用 `make(chan val-type)` 创建一个新的通道。
}
