package main

import (
  "net"
  "fmt"
  "bufio"
  "os"
)

func main() {
  runningClient := launchClient(os.Args[1], os.Args[2])
  userMessage(runningClient)
}
// In the future, make the arguments, host String, port String, file File
func launchClient(host, port string) net.Conn{
  conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
  return conn
}

// Sends a simple message to the server, receives a response
func userMessage(ln net.Conn) {
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    fmt.Fprintf(ln, text + "\n")
    message, _ := bufio.NewReader(ln).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}
