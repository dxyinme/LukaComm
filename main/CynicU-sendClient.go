package main

import (
	"github.com/dxyinme/LukaComm/CynicU/SendMsg"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"sync"
	"time"
)

func main() {
	msg := &chatMsg.Msg{
		From:           "1",
		Target:         "2",
		Content:        []byte("say hi!!!!!!!!!!!!!!!!!!!!!1"),
		MsgType:        chatMsg.MsgType_Single,
		MsgContentType: chatMsg.MsgContentType_Text,
		SendTime:       "",
		GroupName:      "",
		Spread:         false,
		MsgId:          "",
	}
	reqCount := 8000
	timeoutCnt := 0
	startTime := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(reqCount)
	for i := 0 ; i < reqCount; i ++{
		go func(v int) {
			c := SendMsg.NewClient("127.0.0.1:8080")
			msgNow := *msg
			msgNow.SendTime = time.Now().String()
			err := c.SendTo(&msgNow)
			if err != nil {
				log.Println(err)
			}
			if err == SendMsg.TimeoutErr {
				timeoutCnt ++
			}

			wg.Done()
		}(i)
	}
	wg.Wait()

	log.Printf("use time : %v ms in %v requests, %v timeout",
		time.Now().Sub(startTime).Milliseconds(), reqCount, timeoutCnt)


}
