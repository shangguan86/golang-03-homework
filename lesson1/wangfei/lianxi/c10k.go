package main

import (
	"net"
	"fmt"
	"time"
	"log"
)

func handle(conn net.Conn) {
	fmt.Println(conn, "%s", time.Now().String())
	conn.Close()
}

func main() {
	log.Println("Start to listen on tcp port 8080")
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for {
		conn, err := l.Accept()
		count += 1
		fmt.Println(count)
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

