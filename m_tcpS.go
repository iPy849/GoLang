package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
)

var (
	firstConnection  bool = true
	activeConnection int
)

func HandleConnections(conn *net.Conn, cancelCtx *context.CancelFunc) {
	defer (*conn).Close()
	defer func() { activeConnection-- }()

	for {
		connReader := bufio.NewReader(*conn)
		data, err := connReader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error: " + err.Error())
				(*cancelCtx)()
			}
			return
		}
		add := (*conn).LocalAddr().String()
		fmt.Printf("(%s) => %s", add, data)

		response := fmt.Sprintf("You said: %s", data)
		_, err = (*conn).Write([]byte(response))
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error: " + err.Error())
				(*cancelCtx)()
			}
			return
		}
		fmt.Printf("(%s) <= %s", add, response)
	}
}

func main() {
	port := ":1234"
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	//NOTE: Cada conexión levanta una gorrutina distinta, será buena tenerlas bajo un contexto.
	incomingRequestCtx, cancel := context.WithCancel(context.Background())

	go func() {
		for firstConnection || activeConnection > 0 {
		}
		os.Exit(1)
	}()

	for {
		fmt.Println(firstConnection)
		fmt.Printf("%dC - Ready to receive connections\n", activeConnection)
		// Compruebo que todos se conecten correctamente, o apago el servidor
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: " + err.Error())
			cancel()
		}

		// Si la funciónn terminó toca avisar y hacer limpieza de recursos
		select {
		case <-incomingRequestCtx.Done():
			fmt.Println("Server terminated succesfully!!!")
			runtime.GC() //NOTE: Garbage collector
			return
		default:
			if firstConnection {
				firstConnection = false
			}
			activeConnection++
		}

		fmt.Printf("New connection from %s\n", conn.LocalAddr().String())
		go HandleConnections(&conn, &cancel)
	}

}
