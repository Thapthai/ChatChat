package controllers

import (
	"ChatChat/models"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // ปรับให้เหมาะสมกับความต้องการความปลอดภัย
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		var msg models.Message

		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading json:", err)
			break
		}

		fmt.Printf("Received message from %s: %s\n", msg.Username, msg.Content)

		if err := conn.WriteJSON(msg); err != nil {
			fmt.Println("Error writing json:", err)
			break
		}
	}
}
