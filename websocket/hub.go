package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Hub struct {
	Name       string
	Clients    map[*Client]Vote
	IsUnveiled bool
	//broadcast  chan []byte
	//register   chan *Client
	//unregister chan *Client
	terminate chan<- *Hub
	commands  chan command
}

func newHub(name string, terminate chan *Hub) *Hub {
	return &Hub{
		Name:       name,
		Clients:    make(map[*Client]Vote),
		IsUnveiled: false,
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
			fmt.Println("join the room:", cmd.client)
			h.Clients[cmd.client] = Vote{HasValue: false}
		case CMD_List:
		case CMD_UNVEIL:
			h.IsUnveiled = true
		case CMD_CLEAR:
			// veil the room and reset votes
			h.IsUnveiled = false
			for k := range h.Clients {
				h.Clients[k] = Vote{HasValue: false}
			}
		case CMD_QUIT:
			if _, ok := h.Clients[cmd.client]; ok {
				log.Println("unresiger: removing client...", cmd.client.nick)
				delete(h.Clients, cmd.client)
				close(cmd.client.send)
			}
		case CMD_NICK:
			if len(cmd.args) < 2 {
				log.Println("invalid rename command:", cmd.args)
				continue
			}
			log.Printf("renaming client: %s -> %s\n", cmd.client.nick, cmd.args[1])
			cmd.client.nick = cmd.args[1]
		case CMD_TERM:
			break
		case CMD_VOTE:
			log.Println(cmd.args)
			if len(cmd.args) < 2 {
				log.Println("invalid rename command:", cmd.args)
				continue
			}
			//cmd.client.hub.broadcast()
			vote, err := strconv.Atoi(cmd.args[1])
			if err != nil {
				fmt.Println("Invalid vote:", cmd.args[1])
			}
			h.Clients[cmd.client] = Vote{HasValue: true, V: vote}
			break
		default:
			log.Println("Hub unhandled command:", cmd)
		}

		// broadcast the room to all clients
		roomInfo, _ := json.Marshal(Response{Type: "room", Value: h.buildRoom()})
		h.broadcast(string(roomInfo))

		if len(h.Clients) == 0 {
			go func() {
				timer := time.NewTimer(time.Second * 600)
				log.Printf("waiting for %d seconds.\n", 10)
				<-timer.C
				log.Printf("Total %d clients after %d seconds.", len(h.Clients), 600)
				if len(h.Clients) == 0 {
					h.commands <- command{id: CMD_TERM}
					terminate <- h
				}
			}()
		}
	}
}

type ClientVote struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
	Vote Vote   `json:"vote"`
}

func (h *Hub) buildVotes() []ClientVote {
	list := make([]ClientVote, 0)
	for k, v := range h.Clients {
		list = append(list, ClientVote{ID: k.id, Nick: k.nick, Vote: v})
	}

	return list
}

// define type for websocket response
// it has type that indicates the type of message
// and a value field which contains the data
type Response struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Vote struct {
	HasValue bool `json:"hasValue"`
	V        int  `json:"v"`
}

type Room struct {
	Name       string       `json:"name"`
	Votes      []ClientVote `json:"clients"`
	IsUnveiled bool         `json:"isUnveiled"`
}

func (h *Hub) buildRoom() Room {
	return Room{
		Name:       h.Name,
		Votes:      h.buildVotes(),
		IsUnveiled: h.IsUnveiled,
	}
}

func (h *Hub) broadcast(message string) {
	for client := range h.Clients {
		select {
		case client.send <- []byte(message):
		default:
			log.Println("send: removing client...", client.nick)
			close(client.send)
			delete(h.Clients, client)
		}
	}
}
