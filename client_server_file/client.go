package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  data, err := os.ReadFile("text.txt")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  conn, err := net.Dial("tcp", "localhost:80")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer conn.Close()

  _, err = conn.Write([]byte(data))
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

}
