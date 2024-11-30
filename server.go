package main

import (
	"fmt"
	"net"
	"os"
)

func main(){
	listener, listenerErr := net.Listen("tcp", "0.0.0.0:8379")
	fmt.Println("Server listening on port 0.0.0.0:8379")

	if listenerErr != nil {
		os.Exit(1)
	}
	// defer listener.Close()

	connection, connectionErr := listener.Accept()

	if connectionErr != nil {
		if connectionErr.Error() == "EOF" {
			fmt.Println("no more file to read")
			connection.Close()
			os.Exit(1)
		}
	}

	buf := make([]byte, 1024)
	for {
		connection.Read(buf)
		message := "Return: " + string(buf)
		fmt.Println(message)
		connection.Write([]byte(message))
	}
		
}