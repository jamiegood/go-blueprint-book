package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// client ...
type client struct {
	socket   *websocket.Conn
	fromRoom chan []byte
	room     *room
}

func (c *client) sendToRoom() {
	defer c.socket.Close()
	_, msg, err := c.socket.ReadMessage()

	if err != nil {
		return
	}

	fmt.Println("SENDING messgae to room quque for clients ")
	c.room.sendmsgtoclients <- msg
}

func (c *client) readFromRoom() {
	defer c.socket.Close()

	//for msg := range c.fromRoom {
	for msg := range c.room.sendmsgtoclients {
		c.socket.WriteMessage(websocket.TextMessage, msg)
		fmt.Println("sent message to the client via the wesocket")

	}

}
