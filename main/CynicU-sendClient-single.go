package main

import (
	"github.com/dxyinme/LukaComm/CynicU/SendMsg"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"sync"
	"time"
)

func main() {
	Cnt := 1000
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

	startTime := time.Now()
	timeout := 0
	c := SendMsg.NewClient("127.0.0.1:8080")
	wg := sync.WaitGroup{}
	wg.Add(Cnt)
	for i := 0 ; i < Cnt ; i ++ {
		go func() {
			msgNow := *msg
			msgNow.SendTime = time.Now().String()
			err := c.SendTo(&msgNow)
			if err != nil {
				log.Println(err)
				timeout ++
			}
			wg.Done()
		}()
		//time.Sleep(5 * time.Millisecond)
	}
	wg.Wait()

	log.Println("time use:", time.Now().Sub(startTime))
	log.Println("time out:", timeout)

}