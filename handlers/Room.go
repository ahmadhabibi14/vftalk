package handlers

import (
	"sync"
	"vftalk/conf"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
)

type (
	MessageIn struct {
		Message string `json:"message"`
	}
	MessageOut struct {
		Username string `json:"username"`
		Message  string `json:"message"`
	}
)

// Add more data to this type if needed
type client struct {
	isClosing bool
	mu        sync.Mutex
}

var clients = make(map[*websocket.Conn]*client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var register = make(chan *websocket.Conn)
var broadcast = make(chan MessageOut)
var unregister = make(chan *websocket.Conn)

func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = &client{}

		case message := <-broadcast:
			// Send the message to all clients
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) { // send to each client in parallel so we don't block on a slow client
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}
					if err := connection.WriteJSON(message); err != nil {
						c.isClosing = true

						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-unregister:
			// Remove the client from the hub
			delete(clients, connection)
		}
	}
}

func HandleClients(conn *websocket.Conn) {
	go runHub()
	// When the function returns, unregister the client and close the connection
	defer func() {
		unregister <- conn
		conn.Close()
	}()

	// Register the client
	register <- conn

	username, _ := conf.WsGetUsernameFromJWT(conn)
	for {
		var messageIn MessageIn
		err := conn.ReadJSON(&messageIn)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("Read error:", err)
			}
			return // Calls the deferred function, i.e. closes the connection on error
		}

		messageOut := MessageOut{
			Username: username.(string),
			Message:  messageIn.Message,
		}
		broadcast <- messageOut
	}
}
