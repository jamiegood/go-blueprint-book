package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

// Room ...
type room struct {
	name             string
	clients          []client
	sendmsgtoclients chan []byte
}

// newRoom create
func newRoom() *room {
	fmt.Println("in room constructor")
	return &room{name: "jamiewashere"}
}

// ServeHTTP handles the HTTP request.
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := client{
		socket:   socket,
		fromRoom: make(chan []byte, messageBufferSize),
		room:     r,
	}

	r.joinRoom(client)
	client.sendToRoom()
	client.readFromRoom()
	// r.join <- client
	// defer func() { r.leave <- client }()
	//go client.write()
	//client.read()
}

func (r *room) run() {

}

func (r *room) joinRoom(c client) {
	r.clients = append(r.clients, c)
	fmt.Println("addeed a client")
	//
}

// hello
func (r *room) sayHello() {
	fmt.Println("say hello")
}
