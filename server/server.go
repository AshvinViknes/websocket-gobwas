package main

import (
	"log"
	"net"
	"net/http"
	"websocket-gobwas/mongoDB"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var clients = make(map[net.Conn]bool)

func main() {
	mongoDB.SetupMongoDB()

	http.HandleFunc("/chat", HandleChat)
	log.Println("Server listening on :7777")
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func HandleChat(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("New client connected:", conn.RemoteAddr())
	clients[conn] = true
	defer func() {
		delete(clients, conn)
		log.Println("Client disconnected:", conn.RemoteAddr())
	}()

	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received message from %s: %s", conn.RemoteAddr(), msg)

		err = mongoDB.SaveMessageToMongo(string(msg))
		if err != nil {
			log.Println("Failed to save message to MongoDB:", err)
		}
		response := []byte("I received your message")
		err = wsutil.WriteServerMessage(conn, op, response)
		if err != nil {
			log.Println("Failed to send response:", err)
		}
		for client := range clients {
			if client != conn {
				err := wsutil.WriteServerMessage(client, op, msg)
				if err != nil {
					log.Println("Failed to send message:", err)
				}
			}
		}
	}
}
