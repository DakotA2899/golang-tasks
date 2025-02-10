package main

import (
  "fmt"
  "log"
  "net"
  "os"
)

func main() {
  listener, err := net.Listen("tcp", "localhost:80")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println(err)
      continue
    }
    go func() {
      defer conn.Close()
      buffs := make([]byte, 1024)
      n, err := conn.Read(buffs)
      if err != nil {
        fmt.Println(err)
        return
      }

      err = os.WriteFile("server_text.txt", buffs[0:n], 0644)
      if err != nil {
        log.Fatal(err)
      }
      fmt.Println("файл сохранен на сервере")
    }()
  }
}
