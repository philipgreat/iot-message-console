package main

import (
	"fmt"
	"net"
	"os"
)

/* CheckError A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
func decode(b []byte, length int) ([]byte, error) {

	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = 255 - b[i]
	}
	return buf, nil

}
func main() {
	//src := []byte("48656c6c6f20476f7068657221")

	//dst := make([]byte, hex.DecodedLen(len(src)))
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":54321")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		decodedBuffer, err := decode(buf, n)

		fmt.Println("Received ", string(decodedBuffer[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}

	}
}
