package main

import (
  "net"
  "fmt"
  "os"
  "encoding/gob"
  "bytes"
  "bufio"
  "io"
  // "encoding/binary"
  tcp "./tcp"
  "log"
)

func fileToByteArr(input string) []byte {
  // var data []byte
  // var eol bool
  // var str_array []string

  file, err := os.Open(input)
  if err != nil {
    panic(err)
  }

  defer file.Close()

  reader := bufio.NewReader(file)
  buffer := bytes.NewBuffer(make([]byte, 0, 1024))

  for {
    data, _, err := reader.ReadRune();
    if err != nil {
      if err == io.EOF {
        break
      } else {
        panic(err)
      }
    }
    buffer.Write([]byte(string(data)))
  }

  // if err == io.EOF {
  //   err = nil
  // }
  result := []byte(buffer.String())
  buffer.Reset()
  return result
}

func main() {
  // inputFile := []byte(os.Args[1])
  // fileIntoBytes := fileToByteArr(os.Args[1])
  typeOf := "unix" // or "unixgram" or "unixpacket"
  laddr := net.UnixAddr{"/tmp/unixdomaincli", typeOf}
  conn, err := net.DialUnix(typeOf, &laddr/*can be nil*/,
      &net.UnixAddr{"/tmp/unixdomain", typeOf})
  if err != nil {
      panic(err)
  }
  defer os.Remove("/tmp/unixdomaincli")

  var network bytes.Buffer        // Stand-in for a network connection
  enc := gob.NewEncoder(&network) // Will write to network.

  for {
    // buf := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a file name: ")
    userInput := bufio.NewReader(os.Stdin)
    fileName, err := userInput.ReadBytes('\n')
    if err != nil {
      panic(err)
    }
    fileName = fileName[:len(fileName)-1]
    // fmt.Print(fileName)
    // fileNameString = string(fileName)

    fmt.Println()
    // inputFile := []byte(fileName)
    // inputFile := fileName
    fileIntoBytes := fileToByteArr(string(fileName))
    // var network bytes.Buffer        // Stand-in for a network connection
    // enc := gob.NewEncoder(&network) // Will write to network.

    testHeader := tcp.TCPHeader {
      Options: []tcp.TCPOptions {
        tcp.TCPOptions {
          Kind: 0xA0,
          Length: 0xFF,
          Data: fileIntoBytes,
          FileName: fileName,
        },
        tcp.TCPOptions {Kind: 0xAA, Length: 0xAA},
      },
    }
    fmt.Println("First encode")
    testingHeaderEncode := enc.Encode(testHeader)
    if testingHeaderEncode != nil {
        log.Fatal("encode error:", testingHeaderEncode)
    }

    fmt.Println("First write")
    _, err = conn.Write(network.Bytes())
    if err != nil {
        panic(err)
    }

    // conn, err = net.DialUnix(typeOf, &laddr/*can be nil*/,
    //     &net.UnixAddr{"/tmp/unixdomain", typeOf})
    //
    // testHeader.Options[0].Length = 0xAF
    // fmt.Println("Second encode")
    // testingHeaderEncode = enc.Encode(testHeader)
    // if testingHeaderEncode != nil {
    //     log.Fatal("encode error:", testingHeaderEncode)
    // }
    // fmt.Println("Second write")
    // _, err = conn.Write(network.Bytes())
    // if err != nil {
    //     panic(err)
    // }
// end of for loop
  }


  conn.Close()
}
