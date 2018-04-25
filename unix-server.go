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
  fileName := "new" + string(input.Options[0].FileName)
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

// func closeServer() {
//   os.Remove("/tmp/unixdomain")
// }

func main() {
  l, err := net.ListenUnix("unix",  &net.UnixAddr{"/tmp/unixdomain", "unix"})
  fmt.Print("Let's listen\n")
  if err != nil {
     panic(err)
  }
  defer os.Remove("/tmp/unixdomain")

  fmt.Println("BEGINNING MY DUDE")
  conn, err := l.AcceptUnix()
  fmt.Print("Found connection\n")
  if err != nil {
     panic(err)
  }

  x := 0
  for {
    fmt.Println("LOOP NUMBER ", x, " MY DUDE")
    // fmt.Println("beginning")
    // conn, err := l.AcceptUnix()
    // fmt.Print("Found connection\n")
    // if err != nil {
    //    panic(err)
    // }
    // encoder := gob.NewEncoder(conn)
    fmt.Println("ESTABLISH DECODER MY DUDE")
    decoder := gob.NewDecoder(conn)

    // // // Decode (receive) the value.
    fmt.Println("SAMPLE STRUCT MY DUDE")
    var testingHeaderDecode tcp.TCPHeader
    // testingHeaderDecode = tcp.TCPHeader {
    //   Options: []tcp.TCPOptions {
    //     tcp.TCPOptions {Kind: 0x00, Length: 0x00},
    //     tcp.TCPOptions {Kind: 0x00, Length: 0x00},
    //   },
    // }
    fmt.Println("DECODE THAT STRUCT MY DUDE")
    err = decoder.Decode(&testingHeaderDecode)
    // testHeader = decoder.Decode(&q)

    fmt.Println("ERROR CHECKING MY DUDE")
    if err != nil {
        log.Fatal("decode error:", err)
    }
    // fmt.Printf(string(testingHeaderDecode.Options[0].Kind))

    // fmt.Println(testingHeaderDecode.Options[0].Data)
    fmt.Println("MAKE IT A FILE MY DUDE")
    byteArrToFile(testingHeaderDecode)
    // fmt.Printf("let's seperate these two");
    // fmt.Printf("LIKE REALLY SEPERATE THEM");


    // fmt.Printf("Testing this out: %s\n", string(buf[:n]));


    // fmt.Printf(n);
    // fmt.Printf("Testing this out: %s\n", "sure");
    // fmt.Print("")
    fmt.Println("end")
    x += 1
    // conn.Close()
  }
  conn.Close()
}
