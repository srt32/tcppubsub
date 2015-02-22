package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:4000")
	defer ln.Close()
	if err != nil {
		fmt.Printf("Problem listening: %v", err)
		os.Exit(1)
	}
	fmt.Println("Listening!")

	for {
		fmt.Println("INSIDE OF accept")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Problem accepting: %v", err)
			os.Exit(1)
		}

		fmt.Println("FROM BEFORE handleConnection")
		//go handleConnection(conn)
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("FROM INSIDE handleConnection")

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Problem reading: %v", err.Error())
	}

	fmt.Printf("Sending response for %v bytes \n", n)

	conn.Write([]byte(fmt.Sprintf("Recieved: %v bytes \n", n)))

	conn.Close()
}
