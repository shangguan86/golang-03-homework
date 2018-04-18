package main

import (
	"net"
	"fmt"
	"time"
	"log"
)

func handle(conn net.Conn) {
	fmt.Fprintf(conn, "%s", time.Now().String())
	conn.Close()
}

func main() {
	log.Println("Start to listen on tcp port 8090")
	l, err := net.Listen("tcp", ":8090")
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
