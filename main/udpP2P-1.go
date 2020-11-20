package main

import (
	"github.com/dxyinme/LukaComm/CynicU/P2P"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"strconv"
	"time"
)

func main() {
	p2p := P2P.NewP2P()
	err := p2p.Listen(":8080")
	var now = 0
	for {
		now ++
		time.Sleep(10 * time.Second)
		err = p2p.Send(&chatMsg.Msg{
			From:           "1",
			Target:         "2",
			Content:        []byte("hello" + strconv.Itoa(now)),
			MsgType:        chatMsg.MsgType_Single,
			MsgContentType: chatMsg.MsgContentType_Text,
			SendTime:       "?",
			GroupName:      "-",
			Spread:         false,
		}, "localhost:7070")
		if err != nil {
			log.Println(err)
		}
	}
}
