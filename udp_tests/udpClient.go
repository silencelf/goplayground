package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := "localhost:8008"

	udpAddr, _ := net.ResolveUDPAddr("udp", serverAddr)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	message := []byte("Hello, UDP Server!")

	_, err = conn.Write(message)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Sent message to", udpAddr)
}
