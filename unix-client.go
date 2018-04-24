package main

import (
  "net"
  // "fmt"
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
  inputFile := []byte(os.Args[1])
  fileIntoBytes := fileToByteArr(os.Args[1])
  typeOf := "unix" // or "unixgram" or "unixpacket"
  laddr := net.UnixAddr{"/tmp/unixdomaincli", typeOf}
  conn, err := net.DialUnix(typeOf, &laddr/*can be nil*/,
      &net.UnixAddr{"/tmp/unixdomain", typeOf})
  if err != nil {
      panic(err)
  }
  defer os.Remove("/tmp/unixdomaincli")

  // Initialize the encoder and decoder.  Normally enc and dec would be
  // bound to network connections and the encoder and decoder would
  // run in different processes.
  var network bytes.Buffer        // Stand-in for a network connection
  enc := gob.NewEncoder(&network) // Will write to network.



  // enc := gob.NewEncoder(conn) // Will write to network.
  // dec := gob.NewDecoder(&network) // Will read from network.
  // Encode (send) the value.
  testHeader := tcp.TCPHeader {
    Options: []tcp.TCPOptions {
      tcp.TCPOptions {
        Kind: 0xA0,
        Length: 0xFF,
        Data: fileIntoBytes,
        FileName: inputFile,
      },
      tcp.TCPOptions {Kind: 0xAA, Length: 0xAA},
    },
  }
  testingHeaderEncode := enc.Encode(testHeader)
  if testingHeaderEncode != nil {
      log.Fatal("encode error:", testingHeaderEncode)
  }
  // fmt.Println(network.Bytes())



  // fmt.Println(testingHeaderEncode)
  // Decode (receive) the value.
  // var q Q
  // err = dec.Decode(&q)
  // if err != nil {
  //     log.Fatal("decode error:", err)
  // }
  // fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)



  _, err = conn.Write(network.Bytes())
  // _, err = conn.Write([]byte("hello"))

  // _, err = conn.Write([]byte(tcp.TCPHeader {
  //   Options: []tcp.TCPOptions {
  //     tcp.TCPOptions {Kind: 0xFF, Length: 0xFF},
  //     tcp.TCPOptions {Kind: 0xAA, Length: 0xAA},
  //   },
  // }))
  if err != nil {
      panic(err)
  }
  conn.Close()
}
