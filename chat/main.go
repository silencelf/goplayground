package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	port := ":8888"
	go s.run()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("started server on %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connect: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
