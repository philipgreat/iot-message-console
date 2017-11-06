package main

import (
	"fmt"
	"net"
	"os"
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

	buf := make([]byte, 1024)

	for {
		n, _, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		decodedBuffer, err := decode(buf, n)
		if err != nil {
			fmt.Println("Error: ", err)

		}
		//fmt.Println("Received ", string(decodedBuffer[0:n]), " from ", addr)
		fmt.Println(string(decodedBuffer[0:n]))
		hub.broadcast <- decodedBuffer[0:n]

	}
}
