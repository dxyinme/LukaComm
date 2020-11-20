package main

import (
	"github.com/dxyinme/LukaComm/CynicU/P2P"
	"log"
)

func main() {
	p2p := P2P.NewP2P()
	err := p2p.Listen(":7070")
	if err != nil {
		log.Println(err)
	}
	for {
		select {
		case msg := <-p2p.RecvMsg:
			log.Println(msg)
			break
		}
	}
}