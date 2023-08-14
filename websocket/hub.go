package main

import (
	"log"
	"strings"
)

type Hub struct {
	name    string
	clients map[*Client]bool
	//broadcast  chan []byte
	//register   chan *Client
	//unregister chan *Client
	terminate chan<- *Hub
	commands  chan command
}

func newHub(name string, terminate chan *Hub) *Hub {
	return &Hub{
		name:    name,
		clients: make(map[*Client]bool),
		// broadcast:  make(chan []byte),
		// register:   make(chan *Client),
		// unregister: make(chan *Client),
		commands:  make(chan command),
		terminate: terminate,
	}
}

func (h *Hub) run() {
	for cmd := range h.commands {
		switch cmd.id {
		case CMD_JOIN:
			h.clients[cmd.client] = true
		case CMD_QUIT:
			if _, ok := h.clients[cmd.client]; ok {
				log.Println("unresiger: removing client...", cmd.client.nick)
				delete(h.clients, cmd.client)
				close(cmd.client.send)
			}
		case CMD_NICK:
			if len(cmd.args) < 2 {
				log.Println("invalid rename command:", cmd.args)
				continue
			}
			log.Printf("renaming client: %s -> %s\n", cmd.client.nick, cmd.args[1])
			cmd.client.nick = cmd.args[1]
			h.broadcast(strings.Join(cmd.args, " "), cmd.client)
		default:
			log.Println("Hub unhandled command:", cmd)
		}

		if len(h.clients) == 0 {
			terminate <- h
			break
		}
	}
}

func (h *Hub) broadcast(message string, exclude *Client) {
	for client := range h.clients {
		if client == exclude {
			continue
		}
		select {
		case client.send <- []byte(message):
		default:
			log.Println("send: removing client...", client.nick)
			close(client.send)
			delete(h.clients, client)
		}
	}
}
