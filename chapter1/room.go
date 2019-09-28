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
	//return &room{name: "jamiewashere"}
	return &room{
		name:             "jamiewashere",
		sendmsgtoclients: make(chan []byte),
		//join:    make(chan *client),
		//leave:   make(chan *client),
		clients: make([]client, 10),
	}

}

// ServeHTTP handles the HTTP request.
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		name:     "My client",
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
	fmt.Println("we got to room run")

	for {
		select {
		case msg := <-r.sendmsgtoclients:
			fmt.Println("inside the case")
			fmt.Println(string(msg))

			for _, client := range r.clients {
				client.fromRoom <- msg
				fmt.Println("inside loop")

				//fmt.Println(msg)
				fmt.Println(client.name)
			}

		}
	}

}

func (r *room) joinRoom(c client) {
	r.clients = append(r.clients, c)
	fmt.Println("room:: joinRoom")
	fmt.Println(c.name)
	//
}

// hello
func (r *room) sayHello() {
	fmt.Println("say hello")
}
