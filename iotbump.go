package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf8"
)

func decode(b []byte, length int) ([]byte, error) {

	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = 255 - b[i]
	}
	return buf, nil

}
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
func castIoTMessage(hub *Hub) {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":54321")
	checkError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	checkError(err)
	defer ServerConn.Close()

	buf := make([]byte, 8192)

	for {
		//n, addr, err := ServerConn.ReadFromUDP(buf)
		n, _, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		decodedBuffer, err := decode(buf, n)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		if utf8.Valid(decodedBuffer) {
			hub.broadcast <- decodedBuffer[0:n]
			continue
		}

		fmt.Println("failed to deocde  ", string(buf[0:n]), " from ", *addr)

		if utf8.Valid(buf) {
			hub.broadcast <- buf[0:n]
			continue
		}

		//fmt.Println(addr, "  ", string(decodedBuffer[0:n]))
		//hub.broadcast <- decodedBuffer[0:n]

	}
}
