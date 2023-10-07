# websocket-gobwas
# WebSocket Chat App with MongoDB Storage

This project demonstrates a simple WebSocket-based chat application implemented using the Gobwas WebSocket library in Go. The messages exchanged in the chat are stored persistently in MongoDB for future reference.

## Getting Started

1. **Install Dependencies:**
   Make sure you have Go installed on your system. Additionally, install the required Go packages using the following command:

   ```bash
   go get github.com/gobwas/ws
   go get github.com/gobwas/ws/wsutil
   go get go.mongodb.org/mongo-driver/mongo
   go get go.mongodb.org/mongo-driver/mongo/options
   ```

2. **MongoDB Setup:**
   Set up a MongoDB instance on your local machine and create a database named "chatapp." The application will use a collection named "ChatHistory" to store chat messages.

## Server Implementation (`server.go`)

- The server listens on port `7777` for WebSocket connections.
- WebSocket connections are upgraded using the Gobwas WebSocket library.
- Connected clients are tracked in the `clients` map.
- Messages received from clients are stored in MongoDB using the `mongoDB` package.
- Upon receiving a message, the server sends an acknowledgment to the sender and broadcasts the message to all other connected clients.

## Client Implementation (`client.go`)

- The client connects to the server using the WebSocket protocol.
- Users are prompted to enter a username.
- Messages entered by the user are sent to the server and displayed locally.
- Messages received from the server are displayed on the client.

## MongoDB Integration (`mongoDB.go`)

- The `SetupMongoDB` function initializes a connection to the MongoDB database and creates a collection named "ChatHistory."
- The `SaveMessageToMongo` function stores chat messages in MongoDB along with a timestamp.


## To Run the Application
- Run the server first using the following command.

   ```bash
  go run server/server.go
   ```
- Then run the client using the following command.
  ```bash
  go run client/client.go
   ```
## Note

- Ensure that MongoDB is running before starting the application.
- Update the MongoDB connection URI in `mongoDB.SetupMongoDB` if your MongoDB instance is not running on the default `localhost:27017`.

Feel free to explore and enhance this chat application as needed.


