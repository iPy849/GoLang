package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

var (
	firstConnection  bool = true
	activeConnection int
	waitGroup        sync.WaitGroup
)

// Atiende las conexiones entrantes y les asigna un handler
func HandleIncomingConnections(listener *net.Listener) {
	//NOTE: Cada conexión levanta una gorrutina distinta, será buena tenerlas bajo un contexto.
	for {
		fmt.Println(firstConnection)
		fmt.Printf("%dC - Ready to receive connections\n", activeConnection)
		// Compruebo que todos se conecten correctamente, o apago el servidor
		conn, err := (*listener).Accept()
		if err != nil { // NOTE: Cuando se cierre el listener puede dispararse el Accept con error por concurrencia
			return
		}

		fmt.Printf("New connection from %s\n", conn.LocalAddr().String())
		go HandleConnections(&conn)

		if firstConnection {
			firstConnection = false
		}
		activeConnection++
	}
}

// Handler para cada conexión
func HandleConnections(conn *net.Conn) {
	defer (*conn).Close()
	defer func() { activeConnection-- }()

	for {
		connReader := bufio.NewReader(*conn)
		data, err := connReader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error: " + err.Error())
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
		fmt.Println(1)
		panic(err)
	}
	defer listener.Close()
	waitGroup.Add(1)

	go HandleIncomingConnections(&listener)

	// NOTE: Comprueba las conexiones para saber cuando terminar de escuchar
	go func() {
		for firstConnection || activeConnection > 0 {
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("About to end...")
}
