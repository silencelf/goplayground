package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var count = 0

func handleConnection(c net.Conn) {
	fmt.Println(".")
	for {
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
	}
	c.Close()
}

func main() {
	arguments := os.Args
	fmt.Println(arguments)
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
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
