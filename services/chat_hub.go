package services

import (
	"encoding/json"
	"real-estate-api/config"
	"real-estate-api/models"
	"sync"

	"github.com/gorilla/websocket"
)

// Client Connection
type Client struct {
	UserID uint
	Conn   *websocket.Conn
}

type ChatHub struct {
	Clients    map[uint]*websocket.Conn
	Register   chan *Client
	Unregister chan uint
	Mu         sync.Mutex
}

var GlobalHub = ChatHub{
	Clients:    make(map[uint]*websocket.Conn),
	Register:   make(chan *Client),
	Unregister: make(chan uint),
}

func (h *ChatHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mu.Lock()
			h.Clients[client.UserID] = client.Conn
			h.Mu.Unlock()
		case userID := <-h.Unregister:
			h.Mu.Lock()
			if conn, ok := h.Clients[userID]; ok {
				conn.Close()
				delete(h.Clients, userID)
			}
			h.Mu.Unlock()
		}
	}
}

type WSMessage struct {
	ReceiverID uint   `json:"receiver_id"`
	Message    string `json:"message"`
}

func (h *ChatHub) BroadcastMessage(senderID uint, rawMessage []byte) {
	var msg WSMessage
	if err := json.Unmarshal(rawMessage, &msg); err != nil {
		return
	}

	dbMsg := models.ChatMessage{
		SenderID:   senderID,
		ReceiverID: msg.ReceiverID,
		Message:    msg.Message,
	}
	config.DB.Create(&dbMsg)

	h.Mu.Lock()
	receiverConn, receiverOnline := h.Clients[msg.ReceiverID]
	senderConn, senderOnline := h.Clients[senderID]
	h.Mu.Unlock()

	if receiverOnline {
		receiverConn.WriteJSON(dbMsg)
	}

	if senderOnline {
		senderConn.WriteJSON(dbMsg)
	}

	go SendNotification(
		msg.ReceiverID,
		"New Message",
		"A new letter has arrived.",
		"chat",
	)
}
