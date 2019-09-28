package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// client ...
type client struct {
	name     string
	socket   *websocket.Conn
	fromRoom chan []byte
	//	send chan []byte

	room *room
}

func (c *client) sendToRoom() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(msg)
		fmt.Println("SENDING messgae to room quque for clients ")
		c.room.sendmsgtoclients <- msg
	}

}

// func (c *client) readFromRoom() {
// 	defer c.socket.Close()

// 	fmt.Println("XxxxxXXxxxxxxxxxxxxxxxx     Inside the client::readFromRoom")

// 	for {
// 		msg := <-c.fromRoom
// 		//c.socket.WriteMessage(websocket.TextMessage, msg)
// 		fmt.Println("sent message to the client via the wesocket")
// 		fmt.Println(string(msg))
// 		err := c.socket.WriteMessage(websocket.TextMessage, msg)
// 		if err != nil {
// 			return
// 		}

// 	}

// }

func (c *client) readFromRoom() {
	defer c.socket.Close()
	fmt.Println("inside CLient WRITE")
	for msg := range c.fromRoom {
		fmt.Println(string(msg))
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
