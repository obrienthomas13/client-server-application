package main

import (
  "net"
  "fmt"
  "os"
  "encoding/gob"
  "bytes"
  // "encoding/binary"
  tcp "./tcp"
  "log"
	// "crypto/sha1"
)

// type P struct {
//     X, Y, Z int
//     Name    string
// }

func main() {
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
      tcp.TCPOptions {Kind: 0xFF, Length: 0xFF},
      tcp.TCPOptions {Kind: 0xAA, Length: 0xAA},
    },
  }
  testingHeaderEncode := enc.Encode(testHeader)
  if testingHeaderEncode != nil {
      log.Fatal("encode error:", testingHeaderEncode)
  }
  fmt.Println(network.Bytes())



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
