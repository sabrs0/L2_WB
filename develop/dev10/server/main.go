package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	fmt.Println("Connection accepted from ", c.LocalAddr())
	defer func() {
		fmt.Println("Disconnected : ", c.RemoteAddr())
		c.Close()

	}()
	for {
		c.Write([]byte("What is your name ?\n"))
		r := bufio.NewReader(c)
		clientName, err := r.ReadString('\n')
		if err != nil {
			return
		}
		c.Write([]byte("Hello " + clientName))
	}

	//return

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Start listening")
	for {
		c, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(c)
	}
}
