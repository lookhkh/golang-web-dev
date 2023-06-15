package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	io.WriteString(conn, "\nHello from TCP server\n")
	fmt.Fprintln(conn, "How is your day?")
	fmt.Fprintf(conn, "%v [%v]", "Well, I hope!", conn.RemoteAddr())
	defer conn.Close()

	buf := make([]byte, 0, 4096)
	for {
		_, err := conn.Read(buf)

		if err != nil {
			log.Println(err, conn)
			break
		}

		str := string(buf)

		fmt.Fprintf(conn, "you said %v", str)

		if str == "bye" {
			break
		}
	}

	log.Println(conn, "is closed")
}
