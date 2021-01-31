package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main () {
	tcpServer, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer tcpServer.Close()

	for {
		connection, err := tcpServer.Accept();
		if err != nil {
			log.Panicln(err)
		}
		io.WriteString(connection, "\nHello GoLang\n")
		fmt.Println("got a request")
		connection.Close()
	}
}