package main
import "fmt"
import "time"
import "sync/atomic"
import "runtime"

// use sync/atomic

func main() {
  var ops uint64 = 0

  for i := 0; i < 50; i++ {
    go func(){
      for {
        // AddUint64 atomically adds delta to *addr and returns the new value
        atomic.AddUint64(&ops, 1)
        // yields the processor, allowing other goroutines to run.
        runtime.Gosched()
      }
    }()
  }

  time.Sleep(time.Second)
  // LoadUint64 atomically loads *addr.
  ops_final := atomic.LoadUint64(&ops)
  fmt.Println("ops: ", ops_final)
}