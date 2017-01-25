package main

import "time"
import "fmt"

func main(){
  requests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests <- i
  }
  close(requests)

  limiter := time.Tick(time.Millisecond * 200) //every 200 ms chan

  for req := range requests {
    <-limiter //block
    fmt.Println("Request: ", req, time.Now());
  }

  burstyLimiter := make(chan time.Time, 3) // 3 chan like buffers

  for i := 0; i < 3; i++ {
    burstyLimiter <- time.Now()
  }

  go func(){
    for t := range time.Tick(time.Millisecond * 200) {
      burstyLimiter <- t
    }
  }()
  // 5 request to 3 limited
  burstyRequests :=make(chan int, 5)

  for i := 1; i <= 5; i++ {
    burstyRequests <- i
  }
  close(burstyRequests)

  for req := range burstyRequests {
    <-burstyLimiter
    fmt.Println("Request: ", req, time.Now());
  }
}