package main

import (
	"flag"
	CynicUClient "github.com/dxyinme/LukaComm/CynicU/Client"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"sync"
	"time"
)

var (
	isSleep = flag.Bool("isSleep",false,"wait after send")
)

func main() {
	flag.Parse()
	var err error
	client := &CynicUClient.Client{}
	err = client.Initial("localhost:8080", time.Second * 3)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	err = client.SendTo(&chatMsg.Msg{
		From:           "example",
		Target:         "example2",
		Content:        []byte("test"),
		MsgType:        chatMsg.MsgType_Single,
		MsgContentType: chatMsg.MsgContentType_Text,
		SendTime:       time.Now().String(),
	})
	if err != nil {
		log.Fatalln(err)
	}
	if *isSleep {
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}
}
