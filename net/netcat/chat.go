package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	globalMessages  = make(chan string)
)

var (
	host = flag.String("h", "localhost", "Host")
	port = flag.Int("p", 3090, "Port")
)

// Client 1 -> Server -> HandleConnection(Client1)
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn, message)

	clientName := conn.RemoteAddr().String()
	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)
	globalMessages <- fmt.Sprintf("New client is here, name %s\n", clientName)
	incomingClients <- message //enviamos el canal para estar pendiente

	//Read everything from terminal
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		globalMessages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- message
	globalMessages <- fmt.Sprintf("%s said goodbye\n", clientName)
}
func MessageWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}
