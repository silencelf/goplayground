package main

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_List
	CMD_JOIN
	CMD_VOTE
	CMD_SHOW
	CMD_CLEAR
	CMD_QUIT
	CMD_TERM
)

type command struct {
	id     commandID
	client *Client
	args   []string
}
