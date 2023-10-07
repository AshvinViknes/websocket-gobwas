package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://localhost:7777/chat")
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	log.Print("Enter your username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Failed to read username:", err)
		return
	}

	go ReadMessages(conn)

	for {
		log.Print("Enter your message: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read message:", err)
			return
		}
		message := username + ": " + text
		err = wsutil.WriteClientMessage(conn, ws.OpText, []byte(message))
		if err != nil {
			log.Println("Failed to send message:", err)
			return
		}
	}
}

func ReadMessages(conn net.Conn) {
	for {
		msg, op, err := wsutil.ReadServerData(conn)
		if err != nil {
			log.Println("Failed to read message:", err)
			os.Exit(1)
		}
		if op == ws.OpText {
			log.Println("Received:", string(msg))
		}
	}
}
