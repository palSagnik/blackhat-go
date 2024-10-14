package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	port = 20080
	host = "0.0.0.0"
)

func echo(conn net.Conn) {
	defer conn.Close()

	// Receiving and writing using copy
	if _, err := io.Copy(conn, conn); err != nil {
		log.Println("unable to read/write")
	}
}

func bindPort(port int) (net.Listener, error){
	portString := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", portString)
	if err != nil {
		return nil, err
	}
	
	fmt.Println("port bind successful")
	return listener, nil
}

func main() {
	// binding tcp port
	listener, err := bindPort(port)
	if err != nil {
		log.Fatal("unable to bind port")
	}
	log.Printf("listening on %s:%d", host, port)

	// Accept Loop
	for {
		conn, err := listener.Accept()
		log.Println("accepted connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}

		go echo(conn)
	}	
}