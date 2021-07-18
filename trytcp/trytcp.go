package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

var (
	commands (chan command)
)

type command struct {
	conn net.Conn
	id   string
	args []string
}

func main() {
	listener, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
	defer listener.Close()
	commands = make(chan command)
	go handleCommands()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept client: %s", err.Error())
		}

		go handleConn(conn)
	}
}

func handleCommands() {
	for c := range commands {
		switch c.id {
		case "/nick":
			//fmt.Fprintf(c.conn, "hello %s", "there!")
			c.conn.Write([]byte("> " + "hello!"))
		default:
			c.conn.Write([]byte("> working...\n"))
		}
	}
}

type client struct {
	conn     net.Conn
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Printf("failed to read input: %s", err.Error())
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		c.commands <- command{
			conn: c.conn,
			id:   cmd,
			args: args,
		}
	}
}

func handleConn(conn net.Conn) {
	log.Printf("conn: %#v", conn)
	c := &client{
		conn:     conn,
		commands: commands,
	}
	c.readInput()
}
