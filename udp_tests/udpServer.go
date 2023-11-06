package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := "localhost:8008"

	udpAddr, _ := net.ResolveUDPAddr("udp", serverAddr)

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Received %s from %s\n", buffer[:n], addr)
	}
}
