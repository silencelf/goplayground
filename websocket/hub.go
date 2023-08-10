package main

import "log"

type Hub struct {
	name       string
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	terminate  chan<- *Hub
}

func newHub(name string, terminate chan *Hub) *Hub {
	return &Hub{
		name:       name,
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		terminate:  terminate,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Println("unresiger: removing client...", client.name)
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					log.Println("send: removing client...", client.name)
					close(client.send)
					delete(h.clients, client)
				}
			}
			break
		}

		if len(h.clients) == 0 {
			terminate <- h
			break
		}
	}
}
