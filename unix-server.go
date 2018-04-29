package main

import (
  "net"
  "fmt"
  "os"
  "encoding/gob"
  "log"
  tcp "./tcp"
  "image"
  "image/png"
  "strings"
  "bytes"
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

// Currently works for .png files only
func byteArrToImgFile(input tcp.TCPHeader) {
  fmt.Println("Convert the image!!")
  img, _, _ := image.Decode(bytes.NewReader(input.Options[0].Data))
  fileName := "new" + string(input.Options[0].FileName)
  newFile, err := os.Create(fileName)
  if err != nil {
    panic(err)
  }
  err = png.Encode(newFile, img)
  if err != nil {
    panic(err)
  }
  defer newFile.Close()

}

func byteArrToTxtFile(input tcp.TCPHeader) {
  var fileName string
  diffDir := strings.LastIndex(string(input.Options[0].FileName), "/")
  if diffDir > -1 {
    fileName = string(input.Options[0].FileName)[:diffDir+1] + "new" + string(input.Options[0].FileName)[diffDir+1:]
  } else {
    fileName = "new" + string(input.Options[0].FileName)
  }
  // fileName := "tcp/newtcp.go"
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
}

func main() {
  // hard coding for now to make testing smoother
  unixAddress := "/tmp/unixdomain"
  // unixAddress := os.Args[1]
  l, err := net.ListenUnix("unix",  &net.UnixAddr{unixAddress, "unix"})
  if err != nil {
     panic(err)
  }
  fmt.Println("Listening on address: " + unixAddress)
  defer os.Remove(unixAddress)

  conn, err := l.AcceptUnix()
  fmt.Print("Found connection\n")
  if err != nil {
     panic(err)
  }
  for {
    decoder := gob.NewDecoder(conn)
    var incomingTCPHeader tcp.TCPHeader
    err = decoder.Decode(&incomingTCPHeader)
    if err != nil {
        log.Fatal("decode error:", err)
    }
    if checkIfImageType(string(incomingTCPHeader.Options[0].FileName)) {
      byteArrToImgFile(incomingTCPHeader)
    } else {
      byteArrToTxtFile(incomingTCPHeader)
    }
  }
  fmt.Println("Closing server on: " + unixAddress)
  conn.Close()
}
