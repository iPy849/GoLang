package websockets

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gofiber/websocket/v2"
)

var connectionsAlive []*websocket.Conn
var mutex = new(sync.Mutex)

func ChatHandler(c *websocket.Conn) {
	defer c.Close()
	connectionsAlive = append(connectionsAlive, c)
	log.Println("There's a new user in the room")

	// Creates a context for the connection
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "cancel", cancel)
	ctx = context.WithValue(ctx, "conn", c)

	// Notifica conexión al chat
	message := []byte("joined to the chat!")
	broadcastMessageFrom(ctx, websocket.TextMessage, message)

	var (
		msg []byte
		err error
	)
	for {

		// Checks context is done
		select {
		case <-ctx.Done():
			log.Println("Closing connection with " + c.RemoteAddr().String())
			endConnection(ctx)
			return
		default:
		}

		// When client close the connection, server receives an EOF
		if _, msg, err = c.ReadMessage(); err != nil {
			ctx.Value("cancel").(context.CancelFunc)()
			if err != io.EOF {
				log.Println(err)
			}
			continue
		}

		// Solo envía un mensaje si hay algo que enviar o no es un comando
		if !readCommands(msg, ctx) && len(msg) != 0 {
			go broadcastMessageFrom(ctx, websocket.TextMessage, msg)
		}
	}

}

// Envía un mensaje a todos las conexiones del slice connectionsAlive, excepto al emisor
func broadcastMessageFrom(ctx context.Context, msgType int, msg []byte) error {
	c := ctx.Value("conn").(*websocket.Conn)

	for _, conn := range connectionsAlive {
		mutex.TryLock()
		if c == conn {
			continue
		}
		mutex.Unlock()

		add := c.RemoteAddr().String()
		message := fmt.Sprintf("%s -> %s", add, string(msg))
		mutex.TryLock()
		err := conn.WriteMessage(msgType, []byte(message))
		mutex.Unlock()
		if err != nil {
			return BroadcastMessageError(err.Error())
		}
	}
	return nil
}

// Interpreta diferentes comandos escritos por un usuario
func readCommands(incomingMessage []byte, ctx context.Context) bool {
	order := strings.TrimSpace(string(incomingMessage))
	didExecuted := true
	switch order {
	case "END":
		ctx.Value("cancel").(context.CancelFunc)()
	case "SAY HELLO":
		conn := ctx.Value("conn").(*websocket.Conn)
		message := []byte("Hello there!!!")
		conn.WriteMessage(websocket.TextMessage, message)
	default:
		didExecuted = false
	}
	return didExecuted
}

// Elimina la conexión del slice de conexiones vivas y manda un mensaje de control de cierre
func endConnection(ctx context.Context) {
	c := ctx.Value("conn").(*websocket.Conn)

	broadcastMessageFrom(ctx, websocket.TextMessage, []byte("left the chat!"))

	// Envía mensaje de control de cierre, tengo que ver cómo recibir la info
	signalPrefix := os.Getenv("CHAT_SIGNAL_PREFIX")
	closeMessage := signalPrefix + "Closing connection"
	c.WriteMessage(websocket.CloseMessage, []byte(closeMessage))
	for i, conn := range connectionsAlive {
		if c == conn {
			connectionsAlive = append(connectionsAlive[:i], connectionsAlive[i+1:]...)
			break
		}
	}
}
