package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func RunClient() {
	serverAddr := "localhost:8008"
	udpAddr, _ := net.ResolveUDPAddr("udp", serverAddr)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	go func() {
		for {
			buffer := make([]byte, 1024)
			n, addr, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println(err)
			}
			reply := string(buffer[:n])
			fmt.Printf("reply: %v from %v\n", reply, addr)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("[cmd]$ ")
		input, _ := reader.ReadString('\n')
		message := []byte(input)

		_, err := conn.Write(message)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Sent message to %v\n", udpAddr)

	}
}
