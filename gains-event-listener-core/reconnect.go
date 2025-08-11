package main

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
	"time"
)

func reconnect(conn *websocket.Conn, addr string) {
	log.Println("Attempting to reconnect...")
	for i := 1; i <= maxRetries; i++ {
		log.Printf("Retry %d/%d\n", i, maxRetries)
		time.Sleep(retryInterval)

		newConn, _, err := websocket.DefaultDialer.Dial(addr, nil)
		if err == nil {
			log.Println("Reconnected successfully")
			connErr := conn.Close()
			if connErr != nil {
				log.Println("Failed to close connection: ", err)
			}
			conn = newConn
			return
		}
	}
	log.Println("Failed to reconnect after multiple attempts. Exiting...")
	os.Exit(1)
}
