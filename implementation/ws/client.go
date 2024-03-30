package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

var (
	pongWait       = 10 * time.Second
	pingInterval   = (pongWait * 9) / 10
	maxMessageSize = 512
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager

	egress chan Event
}

func NewClient(conn *websocket.Conn, manger *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manger,
		egress:     make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	c.connection.SetReadLimit(int64(maxMessageSize))

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}
	// Configure how to handle Pong responses
	c.connection.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			log.Println(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v\n", err)
			}
			break
		}

		var request Event

		if err = json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling event: %v\n", err)
			//break
		}

		if err = c.manager.routeEvents(request, c); err != nil {
			log.Printf("error routing event: %v\n\n", err)
		}

		// writing data to egress channel to write data concurrently to ws conn
		for singleClient := range c.manager.clients {
			//msg := "server side -> " + string(payload)
			var response Event

			response.Type = "new_message"
			response.Payload = request.Payload

			singleClient.egress <- response
		}

		log.Println(request)
	}
}

func (c *Client) writeMessages() {
	// Create a ticker that triggers a ping at given interval
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		// Graceful close if this triggers a closing
		c.manager.removeClient(c)
	}()

	for {
		select {
		case event, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed:", err)
				}
				return
			}

			message, err := json.Marshal(event)
			if err != nil {
				log.Println(err)
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("failed to send message: %v\n", err)
			}
			log.Println("message sent")
		case <-ticker.C:
			log.Println("ping")
			// Send the Ping
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("write msg: ", err)
				return
			}

		}
	}
}

// pongHandler is used to handle PongMessages for the Client
func (c *Client) pongHandler(_ string) error {
	// Current time + Pong Wait time
	log.Println("pong")
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}
