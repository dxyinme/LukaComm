package main

import (
	"github.com/dxyinme/LukaComm/CynicU/SendMsg"
	"log"
	"os"
	"time"
)

var (
	recvCntSend int64 = 0
)

func recv(server *SendMsg.Server) {
	var bgT string
	for {
		msg := server.GetMsg()
		recvCntSend ++
		log.Println(msg)
		log.Println(recvCntSend)
		if recvCntSend == 1 {
			bgT = msg.SendTime
		}
		if recvCntSend == 10000 {
			log.Println(bgT, time.Now().String())
		}
	}
}

func main() {
	recvCntSend = 0
	log.Printf("PID=%d\n", os.Getpid())
	s := SendMsg.NewServer()
	go recv(s)
	if err := s.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
