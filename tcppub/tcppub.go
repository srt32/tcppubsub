package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	for {
		conn, err := net.Dial("tcp", "localhost:4000")
		defer conn.Close()
		if err != nil {
			fmt.Printf("Problem dialing: %v", err)
			os.Exit(1)
		}

		var input string
		_, err = fmt.Scanf("%v", &input)
		if err != nil {
			fmt.Printf("Problem getting user input: %v", err)
			os.Exit(1)
		}

		sendMessage(conn, input)
	}
}

func sendMessage(conn net.Conn, input string) {
	fmt.Printf("Message is %v", input)

	_, err := fmt.Fprintf(conn, input+"\n")
	if err != nil {
		fmt.Printf("Error writing to connection: %v", err)
		os.Exit(1)
	}

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading response: %v", err)
		os.Exit(1)
	}

	fmt.Println(status)
}
