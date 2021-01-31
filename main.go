package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main () {
	//tcpServer is a Listener interface
	/*
		it has 3 methods
		// accepts waits for and returns the next connection to the listener
		accept
		// closes the listener
		close
		// returns the listeners network address
		addr
	*/
	tcpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer tcpListener.Close()

	for {
		// calling accept returns a connection
		/*
			connect is interface with both Read and Write methods
		*/
		connection, err := tcpListener.Accept();
		if err != nil {
			log.Panicln(err)
		}
		go handleRequest(connection)
		 io.WriteString(connection, "\nHello GoLang\n")
		 fmt.Println("got a request")
		//connection.Close()


	}
}

func handleRequest(conn net.Conn) {
	// hitting this on a browser will display the http request 
	// hitting this with telnet will display
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}