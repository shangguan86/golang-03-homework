package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn) {
	fmt.Fprintf(conn, "%s", time.Now().String())
	conn.Close()
}

func main() {
	log.Println("Start to listen on tcp port 8080")
	l, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}
