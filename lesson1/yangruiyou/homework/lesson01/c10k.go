package  main

import (
	"net"
	"fmt"
	"time"
	"log"
)

func handle(conn net.Conn){
	fmt.Fprint(conn,"%s",time.Now().String())
}


func main(){
	log.Println("start to listen on tcp port 8080")
	l,err := net.Listen("tcp","8080")
	if err!=nil{
		log.Fatal(err)
	}
	for {
		conn,err:=l.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		go handle(conn)
	}
}