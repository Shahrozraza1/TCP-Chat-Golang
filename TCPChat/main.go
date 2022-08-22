package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()
	log.Printf("starting server on :9090")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}
