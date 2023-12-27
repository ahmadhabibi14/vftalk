package apis

import (
	"sync"
	"vftalk/configs"

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

type client struct {
	isClosing bool
	mu        sync.Mutex
}

var (
	clients    = make(map[*websocket.Conn]*client)
	register   = make(chan *websocket.Conn)
	broadcast  = make(chan MessageOut)
	unregister = make(chan *websocket.Conn)
)

func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = &client{}
		case message := <-broadcast:
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) {
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
			delete(clients, connection)
		}
	}
}

func (a *ApisHandler) UserChatRoom(conn *websocket.Conn) {
	go runHub()
	defer func() {
		unregister <- conn
		conn.Close()
	}()

	register <- conn
	username, _ := configs.WsGetUsernameFromJWT(conn)
	for {
		var messageIn MessageIn
		err := conn.ReadJSON(&messageIn)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("Read error:", err)
			}
			return
		}
		messageOut := MessageOut{
			Username: username.(string),
			Message:  messageIn.Message,
		}
		broadcast <- messageOut
	}
}
