package main

import (
  "net"
  "fmt"
  "os"
  "encoding/gob"
  "log"
  tcp "./tcp"
  // "bufio"
  // "bytes"
)

func byteArrToFile(input tcp.TCPHeader) {
  fileName := "new"
  fileName += string(input.Options[0].FileName)
  // fileName := string(input.Options[0].FileName)
  newFile, err := os.Create(fileName)

  if err != nil {
    panic(err)
  }
  defer newFile.Close()
  stringInput := string(input.Options[0].Data)
  _, err = newFile.WriteString(stringInput)
  if err != nil {
    panic(err)
  }
  // writer := bufio.NewWriter(newFile)

  // buffer := make([]byte,1024)
  // for {
  //     // n, err := r.Read(buf)
  //     // if err != nil && err != io.EOF {
  //     //     panic(err)
  //     // }
  //     // if n == 0 {
  //     //     break
  //     // }
  //
  //     // write a chunk
  //     if _, err := writer.Write(buf[:n]); err != nil {
  //         panic(err)
  //     }
  // }

}

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

    // fmt.Println(testingHeaderDecode.Options[0].Data)
    byteArrToFile(testingHeaderDecode)
    // fmt.Printf("let's seperate these two");
    // fmt.Printf("LIKE REALLY SEPERATE THEM");


    // fmt.Printf("Testing this out: %s\n", string(buf[:n]));


    // fmt.Printf(n);
    // fmt.Printf("Testing this out: %s\n", "sure");
    // fmt.Print("")
    conn.Close()
  }
}
