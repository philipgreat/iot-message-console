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
		n, addr, err := ServerConn.ReadFromUDP(buf)
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
			// 在数据前加上IP地址标记
			ipPrefix := []byte("[" + addr.IP.String() + "] ")
			combinedData := make([]byte, len(ipPrefix) + n)
			copy(combinedData, ipPrefix)
			copy(combinedData[len(ipPrefix):], decodedBuffer[0:n])
			
			hub.broadcast <- combinedData
			continue
		}
		
		fmt.Println("failed to decode ", string(buf[0:n]), " from ", addr.IP.String())
	}
}
