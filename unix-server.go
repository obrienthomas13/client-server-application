package main

import (
  "net"
  "fmt"
  "os"
  "encoding/gob"
  "log"
  tcp "./tcp"
  // "bytes"
)

func main() {
  fmt.Print("Let's listen\n")
  l, err := net.ListenUnix("unix",  &net.UnixAddr{"/tmp/unixdomain", "unix"})
  if err != nil {
     panic(err)
  }
  defer os.Remove("/tmp/unixdomain")
  for {
    conn, err := l.AcceptUnix()
    fmt.Print("Found connection\n")
    if err != nil {
       panic(err)
    }
    // encoder := gob.NewEncoder(conn)
    decoder := gob.NewDecoder(conn)

    // var network bytes.Buffer        // Stand-in for a network connection
    // dec := gob.NewDecoder(&network) // Will read from network.


    // var buf [1024]byte
    // n, err := conn.Read(buf[:])
    // // _, err := conn.Read(buf[:])
    // if err != nil {
    //    panic(err)
    // }


    // // // Decode (receive) the value.
    var testingHeaderDecode tcp.TCPHeader
    testingHeaderDecode = tcp.TCPHeader {
      Options: []tcp.TCPOptions {
        tcp.TCPOptions {Kind: 0x00, Length: 0x00},
        tcp.TCPOptions {Kind: 0x00, Length: 0x00},
      },
    }
    err = decoder.Decode(&testingHeaderDecode)
    // testHeader = decoder.Decode(&q)

    if err != nil {
        log.Fatal("decode error:", err)
    }
    // fmt.Printf(string(testingHeaderDecode.Options[0].Kind))
    fmt.Println(testingHeaderDecode.Options[0].Data)
    // fmt.Printf("let's seperate these two");
    // fmt.Printf("LIKE REALLY SEPERATE THEM");


    // fmt.Printf("Testing this out: %s\n", string(buf[:n]));


    // fmt.Printf(n);
    // fmt.Printf("Testing this out: %s\n", "sure");
    // fmt.Print("")
    conn.Close()
  }
}
