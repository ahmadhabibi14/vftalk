package apis

import (
	"fmt"
	"log"
	"sync"
	"time"
	"vftalk/configs"
	"vftalk/utils"

	"github.com/gofiber/contrib/websocket"
	"github.com/rs/zerolog"
)

type GeneralClient struct {
	mu        sync.Mutex
	isClosing bool
}

var (
	GENERAL_CLIENTS    = make(map[*websocket.Conn]*GeneralClient)
	GENERAL_REGISTER   = make(chan *websocket.Conn)
	GENERAL_UNREGISTER = make(chan *websocket.Conn)
	GENERAL_BROADCAST  = make(chan ChatOut)
)

func GeneralBroadcaster(zlog *zerolog.Logger) {
	defer utils.Recover(zlog)
	for {
		select {
		case reg := <-GENERAL_REGISTER:
			GENERAL_CLIENTS[reg] = &GeneralClient{}
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
		case unreg := <-GENERAL_UNREGISTER:
			delete(GENERAL_CLIENTS, unreg)
		}
	}
}

func (a *ApisHandler) ChatRoomGeneral(conn *websocket.Conn) {
	go GeneralBroadcaster(a.Log)
	defer func() {
		GENERAL_UNREGISTER <- conn
		conn.Close()
	}()

	GENERAL_REGISTER <- conn
	username, _ := configs.WsGetUsernameFromJWT(conn)

	msg := fmt.Sprintf("%v join room", username)
	info := ChatOut{
		Type:      CHAT_TYPE_INFO,
		Content:   msg,
		Sender:    CHAT_SENDER_SYSTEM,
		Timestamp: time.Now().UTC(),
	}
	GENERAL_BROADCAST <- info

	for {
		var in ChatIn
		if err := conn.ReadJSON(&in); err != nil {
			msg := fmt.Sprintf("%v sends an invalid message, this incident will be reported", username)
			info := ChatOut{
				Type:      CHAT_TYPE_ERROR,
				Content:   msg,
				Sender:    CHAT_SENDER_SYSTEM,
				Timestamp: time.Now().UTC(),
			}
			GENERAL_BROADCAST <- info
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(err)
			}
			return
		}

		chat := ChatOut{
			Sender:    username.(string),
			Type:      in.Type,
			Content:   in.Content,
			Timestamp: time.Now().UTC().UTC(),
		}

		GENERAL_BROADCAST <- chat
	}
}
