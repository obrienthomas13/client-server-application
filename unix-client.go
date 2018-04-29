package main

import (
  "net"
  "fmt"
  "os"
  "encoding/gob"
  "bytes"
  "bufio"
  "io"
  tcp "./tcp"
  "log"
  "strings"
  "image/png"
)

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
  // return img

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
  laddr := net.UnixAddr{"/tmp/unixdomaincli", typeOf}
  conn, err := net.DialUnix(typeOf, &laddr/*can be nil*/,
      &net.UnixAddr{"/tmp/unixdomain", typeOf})
  if err != nil {
      panic(err)
  }
  defer os.Remove("/tmp/unixdomaincli")

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
    testingHeaderEncode := enc.Encode(testHeader)
    if testingHeaderEncode != nil {
        log.Fatal("encode error:", testingHeaderEncode)
    }
    _, err = conn.Write(payload.Bytes())
    if err != nil {
        panic(err)
    }
  }
  conn.Close()
}
