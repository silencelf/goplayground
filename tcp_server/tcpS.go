package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
)

var count = 0

func handleConnection(c net.Conn) {
	fmt.Println(".")
	for {
		reader := bufio.NewReader(c)
		req, err := http.ReadRequest(reader)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		fmt.Println(req.URL.Path)
		/*
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			if strings.TrimSpace(string(netData)) == "STOP" {
				break
			}

			fmt.Println(c.RemoteAddr(), "-> ", string(netData))
			counter := strconv.Itoa(count) + " "
			t := time.Now()
			resp := counter + t.Format(time.RFC3339) + "\n"
			c.Write([]byte(resp))
		*/
	}
	c.Close()
}

func main() {
	arguments := os.Args
	fmt.Println(arguments)
	port := ":5004"
	if len(arguments) == 2 {
		port = ":" + arguments[1]
	} else {
		fmt.Println("no port provided, using default " + port)
	}

	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
		fmt.Printf("handling %d connections.", count)
	}
}
