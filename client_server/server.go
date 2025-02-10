package main

import (
  "fmt"
  "net"
  "os"
  "time"
)

func main() {
  listener, err := net.Listen("tcp", "localhost:80")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  conn, err := listener.Accept()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  func() {
    _, err = conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
    if err != nil {
      fmt.Println(err)
      return
    }
  }()

  conn.Close()
}
