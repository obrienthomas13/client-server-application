package main

import (
  "bufio"
  "bytes"
  "encoding/gob"
  "fmt"
  handshake "./handshake"
  "os"
  "image/png"
  "io"
  "log"
  "net"
  "strings"
  tcp "./tcp"
)

func initialConnection(conn *net.UnixConn) {
  var payload bytes.Buffer
  enc := gob.NewEncoder(&payload)
  tcpHeaderSetup := tcp.TCPHeader {
    SequenceNumber: 0x0001,
    AcknowledgementNumber: 0x0000,
  }
  outgoingPacket := enc.Encode(tcpHeaderSetup)
  if outgoingPacket != nil {
      log.Fatal("encode error:", outgoingPacket)
  }
  _, err := conn.Write(payload.Bytes())
  if err != nil {
      panic(err)
  }

  decoder := gob.NewDecoder(conn)
  var incomingTCPHeader tcp.TCPHeader
  err = decoder.Decode(&incomingTCPHeader)
  if err != nil {
      log.Fatal("decode error:", err)
  }
  modifiedPacket, result := handshake.ConfirmInitConnection(incomingTCPHeader)
  if !result {
    panic("3-way-handshake failed")
  }
  // Go does not allow a hard wipe on bytes.Buffers.
  // This requires the program to create a new buffer.
  var payloadConfirm bytes.Buffer
  encConfirm := gob.NewEncoder(&payloadConfirm)
  outgoingPacketConfirm := encConfirm.Encode(modifiedPacket)
  if outgoingPacketConfirm != nil {
      log.Fatal("encode error:", outgoingPacketConfirm)
  }
  _, err = conn.Write(payloadConfirm.Bytes())
  if err != nil {
      panic(err)
  }
}

// func listenForResponse(clientAddr string) {
//   fmt.Println("Goroutine: Listen for Response")
//   l, err := net.ListenUnix("unix",  &net.UnixAddr{clientAddr, "unix"})
//   if err != nil {
//      panic(err)
//   }
//   conn, err := l.AcceptUnix()
//   fmt.Print("Found connection\n")
//   if err != nil {
//      panic(err)
//   }
//   decoder := gob.NewDecoder(conn)
//   var incomingTCPHeader tcp.TCPHeader
//   err = decoder.Decode(&incomingTCPHeader)
//   if err != nil {
//       log.Fatal("decode error:", err)
//   }
//   conn.Close()
// }

func checkIfImageType(file string) bool {
  imageTypes := [...]string{".gif", ".jpeg", ".jpg", ".pdf", ".png"}
  for _, imgType := range imageTypes {
    if strings.Contains(file, imgType) {
      return true
    }
  }
  return false
}

func imgFileToByteArr(input string) []byte {
  file, err := os.Open(input)
  defer file.Close()
  if err != nil {
    panic(err)
  }

  img, err := png.Decode(file)
  if err != nil {
    panic(err)
  }

  buffer := new(bytes.Buffer)
  err = png.Encode(buffer, img)
  if err != nil {
    panic(err)
  }
  result := []byte(buffer.String())
  return result
}

func txtFileToByteArr(input string) []byte {
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
  result := []byte(buffer.String())
  buffer.Reset()
  return result
}

func main() {
  var fileIntoBytes []byte
  typeOf := "unix" // or "unixgram" or "unixpacket"
  client := "/tmp/unixdomaincli"
  server := "/tmp/unixdomain"
  laddr := net.UnixAddr{client, typeOf}
  conn, err := net.DialUnix(typeOf, &laddr/*can be nil*/,
      &net.UnixAddr{server, typeOf})
  if err != nil {
      panic(err)
  }
  defer os.Remove(client)

  initialConnection(conn)

  for {
    var payload bytes.Buffer        // Stand-in for a network connection
    enc := gob.NewEncoder(&payload) // Will write to network.
    fmt.Print("Enter a file name: ")
    userInput := bufio.NewReader(os.Stdin)
    fileName, err := userInput.ReadBytes('\n')
    if err != nil {
      panic(err)
    }
    fileName = fileName[:len(fileName)-1]
    if checkIfImageType(string(fileName)) {
      fileIntoBytes = imgFileToByteArr(string(fileName))
    } else {
      fileIntoBytes = txtFileToByteArr(string(fileName))
    }
    tcpHeaderSetup := tcp.TCPHeader {
      Options: []tcp.TCPOptions {
        tcp.TCPOptions {
          Data: fileIntoBytes,
          FileName: fileName,
        },
      },
    }
    outgoingPacket := enc.Encode(tcpHeaderSetup)
    if outgoingPacket != nil {
        log.Fatal("encode error:", outgoingPacket)
    }
    _, err = conn.Write(payload.Bytes())
    if err != nil {
        panic(err)
    }
  }
  conn.Close()
}
