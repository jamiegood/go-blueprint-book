package main

import (
	"fmt"
	"go-blueprint-book/trace"
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
	name string

	sendmsgtoclients chan []byte
	clients          map[*client]bool

	tracer trace.Tracer
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
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
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

	//r.joinRoom(*client)
	r.clients[client] = true

	go client.readFromRoom()

	client.sendToRoom()
	// r.join <- client
	// defer func() { r.leave <- client }()
	//go client.write()
	//client.read()
}

func (r *room) run() {

	for {
		select {
		case msg := <-r.sendmsgtoclients:

			r.tracer.Trace("Message received: ", string(msg))
			fmt.Println(len(r.clients))
			for client := range r.clients {
				//fmt.Println(clientinded)
				r.tracer.Trace(" -- sent to client")

				client.fromRoom <- msg

				//fmt.Println(msg)
				//fmt.Println(client.name)
			}

		}
	}

}

// func (r *room) joinRoom(c client) {
// 	r.clients = append(r.clients, c)
// 	r.tracer.Trace("New client joined")

// 	fmt.Println("room:: joinRoom")
// 	fmt.Println(c.name)
// 	//
// }

// hello
func (r *room) sayHello() {
	fmt.Println("say hello")
}
