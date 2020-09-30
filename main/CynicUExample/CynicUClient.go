package main

import (
	"fmt"
	CynicUClient "github.com/dxyinme/LukaComm/CynicU/Client"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/golang/glog"
	"log"
	"time"
)

func main()  {
	var err error
	var resp *chatMsg.MsgPack
	client := &CynicUClient.Client{}
	err = client.Initial("localhost:8080", time.Second*3)
	if err != nil {
		glog.Error(err)
	}

	err = client.SendTo(&chatMsg.Msg{
		From:           "example",
		Target:         "example2",
		Content:        []byte("test"),
		MsgType:        chatMsg.MsgType_Single,
		MsgContentType: chatMsg.MsgContentType_Text,
	})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = client.PullAll(&chatMsg.PullReq{From: "example2"})
	if err != nil {
		log.Fatalln(err)
	}
	for _, msg := range resp.MsgList {
		fmt.Println(msg)
	}
}