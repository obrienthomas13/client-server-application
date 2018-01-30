package main

import (
  "net"
  "fmt"
  "bufio"
  "strings"
  "os"
)

func main() {
  runningServer := launchServer(os.Args[1])
  listenForClients(runningServer)
}

func launchServer(port string)  net.Conn{
  fmt.Println("Launching server...")
  ln, _ := net.Listen("tcp", ":" + port)

  connection, _ := ln.Accept()
  return connection
}

func listenForClients(ln net.Conn) {
  for {
    message, _ := bufio.NewReader(ln).ReadString('\n')
    fmt.Print("Message Received:", string(message))
    newmessage := strings.ToUpper(message)
    ln.Write([]byte(newmessage + "\n"))
  }
}

func clientEnd() int{
  return 0
}
