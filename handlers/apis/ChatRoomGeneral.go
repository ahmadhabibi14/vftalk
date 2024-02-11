package apis

import (
	"sync"
	"time"
	"vftalk/configs"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
)

type GeneralClient struct {
	isClosing bool
	mu        sync.Mutex
}

var (
	GENERAL_CLIENTS    = make(map[*websocket.Conn]*GeneralClient)
	GENERAL_REGISTER   = make(chan *websocket.Conn)
	GENERAL_UNREGISTER = make(chan *websocket.Conn)
	GENERAL_BROADCAST  = make(chan ChatOut)
)

func (a *ApisHandler) ChatRoomGeneral(conn *websocket.Conn) {
	go func() {
		for {
			select {
			case connection := <-GENERAL_REGISTER:
				GENERAL_CLIENTS[connection] = &GeneralClient{}
			case message := <-GENERAL_BROADCAST:
				for connection, c := range GENERAL_CLIENTS {
					go func(connection *websocket.Conn, c *GeneralClient) {
						c.mu.Lock()
						defer c.mu.Unlock()
						if c.isClosing {
							return
						}
						if err := connection.WriteJSON(message); err != nil {
							c.isClosing = true
							connection.WriteMessage(websocket.CloseMessage, []byte{})
							connection.Close()
							GENERAL_UNREGISTER <- connection
						}
					}(connection, c)
				}
			case connection := <-GENERAL_UNREGISTER:
				delete(GENERAL_CLIENTS, connection)
			}
		}
	}()

	defer func() {
		GENERAL_UNREGISTER <- conn
		conn.Close()
	}()

	GENERAL_REGISTER <- conn
	username, _ := configs.WsGetUsernameFromJWT(conn)
	for {
		var in ChatIn
		err := conn.ReadJSON(&in)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("Read error:", err)
			}
			return
		}
		out := ChatOut{
			Username:  username.(string),
			Type:      in.Type,
			Content:   in.Content,
			Timestamp: time.Now(),
		}
		GENERAL_BROADCAST <- out
	}
}
