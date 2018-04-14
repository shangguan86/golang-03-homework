package main

import (
	"fmt"
	"net"
	"time"
	"log"
)

func handle(conn net.Conn) {
	fmt.Fprintf(conn, "%s", time.Now().String())
	conn.Close()
}

func main() {
	log.Println("Start to listen on tcp port 8888")
	l, err := net.Listen("tcp", "0.0.0.0:8888")
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
