package main

import (
	"github.com/dxyinme/LukaComm/CynicU/SendMsg"
	"log"
)

func recv(server *SendMsg.Server) {
	for {
		msg := server.GetMsg()
		log.Println(msg)
	}
}

func main() {
	s := SendMsg.NewServer()
	go recv(s)
	if err := s.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
