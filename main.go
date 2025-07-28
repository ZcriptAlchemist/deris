package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		PanicOnError(err, "couldn't start the listener")
		return
	}

	log.Println("deris listening on port 6379....")

	conn, err := listener.Accept()
	if err != nil {
		log.Println("couldn't accept connection: ", err)
	}
	defer conn.Close()

	for {

		buffer := make([]byte, 1024)

		if _, err := conn.Read(buffer); err != nil {

			if err == io.EOF {
				log.Println("error reading from the connection", err)
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
		log.Println("received input: ", string(buffer)) // logging user input for devlopment purpose
	}

}
