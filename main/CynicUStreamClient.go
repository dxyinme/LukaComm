package main

import (
	"github.com/dxyinme/LukaComm/CynicU/Client"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"time"
)

func main()  {
	var err error
	var resp *chatMsg.MsgPack
	client := &Client.StreamClient{}
	err = client.Initial("test", "localhost:8080", time.Second*3)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Send(&chatMsg.Msg{
		From:           "tt",
		Target:         "tt",
		Content:        nil,
		MsgType:        0,
		MsgContentType: 0,
		SendTime:       "",
	})
	if err != nil {
		log.Fatal(err)
	}
	resp, err = client.Recv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}