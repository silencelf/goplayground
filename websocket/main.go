package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		runClient()
	}()
	fileServer := http.FileServer(http.Dir("./www"))
	http.Handle("/", fileServer)

	http.Handle("/echo", websocket.Handler(echoHandler))
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Panic("ListenAndServe" + err.Error())
	}
}

func runClient() {
	origin := "http://localhost/"
	url := "ws://localhost:3000/echo"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}
