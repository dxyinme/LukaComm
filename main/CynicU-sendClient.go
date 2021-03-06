package main

import (
	"flag"
	"github.com/dxyinme/LukaComm/CynicU/SendMsg"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"sync"
	"time"
)

var (
	isSleepSend = flag.Bool("isSleep",false,"wait after send")
)

func main() {
	flag.Parse()
	msg := &chatMsg.Msg{
		From:           "example",
		Target:         "example2",
		Content:        []byte("test"),
		MsgType:        chatMsg.MsgType_Single,
		MsgContentType: chatMsg.MsgContentType_Text,
		SendTime:       time.Now().String(),
		GroupName:      "",
		Spread:         false,
		MsgId:          "",
	}

	c := SendMsg.NewClient("127.0.0.1:8080")
	err := c.SendTo(msg)
	if err != nil {
		log.Println(err)
	}
	if *isSleepSend {
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}
}
