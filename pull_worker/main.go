package main

import (
  "fmt"
  "sync"
)

func main() {
  nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

  var wg sync.WaitGroup
  chnums := make(chan int, len(nums))

  for i := range nums {
    chnums <- i
  }
  close(chnums)

  
  for i := 1; i <= 4; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for {
        select {
        case num, ok := <-chnums:
          if !ok {
            return
          }
          res := num * num
          fmt.Println(res)

        }
      }
    }()
  }

  wg.Wait()

}
