package main

import (
  "fmt"
  "errors"
  tcp "./tcp"
)

func main() {
  testHeaderInitial := tcp.TCPHeader {
    Options: []tcp.TCPOptions {
      tcp.TCPOptions {Kind: 0xFF, Length: 0xFF},
      tcp.TCPOptions {Kind: 0xAA, Length: 0xAA},
    },
    SequenceNumber: 0x0001,
    AcknowledgementNumber: 0x0000,
  }
  request, err := initilizeConnection(testHeaderInitial)
  if err != nil {
    panic(err)
  } else {
    fmt.Println(request)
  }
}

func initilizeConnection(packet tcp.TCPHeader) (tcp.TCPHeader, error) {
  acceptRequest := tcp.TCPHeader {
    Options: []tcp.TCPOptions {
      tcp.TCPOptions {Kind: 0xFF, Length: 0xFF},
      tcp.TCPOptions {Kind: 0xAA, Length: 0xAA},
    },
    SequenceNumber: 0x0001,
    AcknowledgementNumber: 0x0001,
  }
  
  if (packet.SequenceNumber == 0x0001 && packet.AcknowledgementNumber == 0x0000) {
    return acceptRequest, nil
  } else {
    return packet, errors.New("Not initial SequenceNumber request")
  }

}
