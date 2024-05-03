package main

import (
	"fmt"
	"net"
	"os"
)


func main() {
	fmt.Println("Starting Fractal Pool...")

	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		fmt.Println("Error listening on port 5000: ", err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Listening on 0.0.0.0:5000")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting tcp conn:", err.Error())
			continue
		}

		go handleTcpConnection(conn)
	}
}

func handleTcpConnection(conn net.Conn) {
	fmt.Println("Received tcp connection")

	conn.Close()
}