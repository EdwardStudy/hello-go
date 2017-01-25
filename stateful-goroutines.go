package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

//使用内置的 Go协程和通道的的同步特性来达到同样的效果 with mutexes.go

type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main(){
  var ops int64

  reads := make(chan *readOp)
  writes := make(chan *writeOp)

  // routine has state
  go func(){
    var state = make(map[int]int)
    
    for {
      //response
      select {
        case read := <- reads:
          read.resp <- state[read.key]
        case write := <- writes:
          state[write.key] = write.val
          write.resp <- true
      }
    }
  }()

  for r := 0; r < 100; r++ {
    //Go 协程通过 reads 通道发起对 state 所有者Go 协程的读取请求
    go func() {
      for {
        read := &readOp{
          key: rand.Intn(5),
          resp: make(chan int),
        }
        reads <- read
        <- read.resp
        atomic.AddInt64(&ops, 1)
      }
    }()
  }

  for w := 0; w < 10; w++ {
    //用相同的方法启动 10 个写操作
    go func() {
      for {
        write := &writeOp{
          key: rand.Intn(5),
          val: rand.Intn(100),
          resp: make(chan bool),
        }
        writes <- write
        <- write.resp
        atomic.AddInt64(&ops, 1)
      }
    }()
  }

  time.Sleep(time.Second)

  opsFinal := atomic.LoadInt64(&ops)
  fmt.Println("ops: ", opsFinal)
}

