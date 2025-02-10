package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  conn, err := net.Dial("tcp", "localhost:80")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer conn.Close()
  buff := make([]byte, 1024)
  n, err := conn.Read(buff)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Println(string(buff[0:n]))

}
