package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"image"
	"image/png"
	"log"
	"net"
	"os"
	"strings"

	handshake "./handshake"
	tcp "./tcp"
)

func confirmHandshake(conn *net.UnixConn, packet tcp.TCPHeader) {
	var payload bytes.Buffer
	enc := gob.NewEncoder(&payload)
	writer := bufio.NewWriter(conn)
	modifiedPacket, result := handshake.InitilizeConnection(packet)
	if result {
		outgoingPacket := enc.Encode(modifiedPacket)
		if outgoingPacket != nil {
			log.Fatal("encode error:", outgoingPacket)
		}
		writer.Write(payload.Bytes())
		writer.Flush()
	}
}

func checkIfImageType(file string) bool {
	imageTypes := [...]string{".gif", ".jpeg", ".jpg", ".pdf", ".png"}
	for _, imgType := range imageTypes {
		if strings.Contains(file, imgType) {
			return true
		}
	}
	return false
}

func parseFileName(input tcp.TCPHeader) (fileName string) {
	diffDir := strings.LastIndex(string(input.Options[0].FileName), "/")
	if diffDir > -1 {
		fileName = "new" + string(input.Options[0].FileName)[diffDir+1:]
	} else {
		fileName = "new" + string(input.Options[0].FileName)
	}
	return
}

// Currently works for .png files only
func byteArrToImgFile(input tcp.TCPHeader) {
	img, _, _ := image.Decode(bytes.NewReader(input.Options[0].Data))
  fileName := parseFileName(input)
	newFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	err = png.Encode(newFile, img)
	defer newFile.Close()

}

func byteArrToTxtFile(input tcp.TCPHeader) {
	fileName := parseFileName(input)
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

func clientHandler(l *net.UnixListener, unixAddress string) {
	conn, err := l.AcceptUnix()
	fmt.Print("Found connection\n")
	if err != nil {
		panic(err)
	}
	go clientHandler(l, unixAddress)
	handShakeOne := false
	handShakeThree := false
	for {
		decoder := gob.NewDecoder(conn)
		var incomingTCPHeader tcp.TCPHeader
		err = decoder.Decode(&incomingTCPHeader)
		if err != nil {
			log.Fatal("decode error:", err)
		}
		if !handShakeOne {
			handShakeOne = true
			confirmHandshake(conn, incomingTCPHeader)
			continue
		} else if !handShakeThree {
			handShakeThree = true
			_, result := handshake.ConfirmPacket(incomingTCPHeader)
			if !result {
				panic("Last stage of 3-way handshake failed")
			}
			fmt.Println("cool it worked!")
			continue
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

func main() {
	unixAddress := os.Args[1]
	l, err := net.ListenUnix("unix", &net.UnixAddr{unixAddress, "unix"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on address: " + unixAddress)
	defer os.Remove(unixAddress)

	clientHandler(l, unixAddress)
}
