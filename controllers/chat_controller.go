package controllers

import (
	"net/http"
	"real-estate-api/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatController struct{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ch ChatController) HandleWS(c *gin.Context) {
	val, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID := val.(uint)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &services.Client{UserID: userID, Conn: conn}
	services.GlobalHub.Register <- client

	defer func() {
		services.GlobalHub.Unregister <- userID
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		services.GlobalHub.BroadcastMessage(userID, message)
	}
}
