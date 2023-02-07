package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	connect := arguments[1]
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">> ")
			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(c, text+"\n")
			if strings.TrimSpace(string(text)) == "STOP" {
				fmt.Println("TCP client exiting...")
				wg.Done()
				return
			}
		}
	}()

	go func() {
		for {
			message, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Printf("->: %s>> ", message)
		}
	}()

	wg.Wait()
}
