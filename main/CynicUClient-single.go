package main

import (
	CynicUClient "github.com/dxyinme/LukaComm/CynicU/Client"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"sync"
	"time"
)

func main() {
	var err error
	client := &CynicUClient.Client{}
	err = client.Initial("localhost:8080", time.Millisecond * 500)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	Cnt := 10000
	startTime := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(Cnt)
	timeout := 0
	for i := 0 ; i < Cnt ; i ++ {
		go func() {
			defer wg.Done()
			err = client.SendTo(&chatMsg.Msg{
				From:           "example",
				Target:         "example2",
				Content:        []byte("test"),
				MsgType:        chatMsg.MsgType_Single,
				MsgContentType: chatMsg.MsgContentType_Text,
				SendTime:       time.Now().String(),
			})
			if err != nil {
				log.Println(err)
				timeout ++
			}
		}()
		//time.Sleep(5 * time.Millisecond)
	}

	wg.Wait()

	log.Println("time use:", time.Now().Sub(startTime))
	log.Println("time out:", timeout)
}
