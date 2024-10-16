package main

import (
	"ChatChat/controllers"
	"ChatChat/database"
	"fmt"
	"net/http"
)

func main() {
	// เรียกใช้ฟังก์ชันเชื่อมต่อฐานข้อมูล
	if err := database.ConnectDatabase(); err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	fmt.Println("Database connection established")

	// ตั้งค่า router สำหรับ WebSocket
	http.HandleFunc("/ws", controllers.HandleWebSocket)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
