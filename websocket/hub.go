package main

import (
	"encoding/json"
	"log"
	"time"
)

const (
	heartbeatTimeout = 600 * time.Second
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
		log.Println(cmd.client.id, cmd.args)
		switch cmd.id {
		case CMD_JOIN:
			log.Println("join the room: ", cmd.client)
			h.Clients[cmd.client] = Vote{HasValue: false}
			// send the client id to the client
			if id, err := json.Marshal(Response{Type: "id", Value: cmd.client.id}); err == nil {
				cmd.client.send <- id
			}
		case CMD_List:
		case CMD_UNVEIL:
			h.IsUnveiled = true
		case CMD_MERGE:
			if len(cmd.args) < 2 {
				log.Println("invalid merge command:", cmd.args)
				continue
			}
			// find the client
			var clients []*Client
			for k := range h.Clients {
				// remove all clients with the same uid
				if k.uid == cmd.args[1] && k.id != cmd.client.id {
					clients = append(clients, k)
				}
			}
			if len(clients) == 0 {
				log.Println("merge: client not found:", cmd.args[1])
				continue
			}
			for _, c := range clients {
				log.Println("merging client:", c.id, c.uid)
				delete(h.Clients, c)
			}
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
			log.Print("exiting the run loop")
		case CMD_VOTE:
			if len(cmd.args) < 2 {
				log.Println("invalid rename command:", cmd.args)
				continue
			}
			// TODO: validate the input
			h.Clients[cmd.client] = Vote{HasValue: true, V: cmd.args[1]}
		default:
			log.Println("Hub unhandled command:", cmd)
		}

		// set last activity time
		if cmd.client != nil {
			cmd.client.lastActivity = time.Now()
		}

		// broadcast the room to all clients
		roomInfo, _ := json.Marshal(Response{Type: "room", Value: h.buildRoom()})
		h.broadcast(string(roomInfo))

		if len(h.Clients) == 0 {
			go func() {
				timeout := 600
				timer := time.NewTimer(time.Second * time.Duration(timeout))
				log.Printf("Hub has no active users. Waiting for %d seconds.\n", timeout)
				<-timer.C
				log.Printf("Total %d clients after %d seconds.", len(h.Clients), timeout)
				if len(h.Clients) == 0 {
					log.Print("terminating...")
					h.commands <- command{id: CMD_TERM}
					log.Print("sending terminate signal...")
					terminate <- h
				}
			}()
		}
	}
}

func (h *Hub) checkHeartbeat() {
	for {
		for k := range h.Clients {
			if time.Since(k.lastActivity) > heartbeatTimeout {
				log.Println("Client heartbeat timeout:", k.nick)
				h.commands <- command{id: CMD_QUIT, client: k}
			}
		}
		time.Sleep(time.Second * 5)
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
	HasValue bool   `json:"hasValue"`
	V        string `json:"v"`
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
