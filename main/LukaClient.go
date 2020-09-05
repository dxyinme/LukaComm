package main

import (
	"github.com/dxyinme/LukaComm/LukaChannel"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/golang/protobuf/proto"
	"log"
)

func main(){
	pod := LukaChannel.Dialer{}
	err := pod.Dial("127.0.0.1:33456", "xxx","xxx2","luka1",10,3)
	if err != nil {
		log.Println(err)
	}
	m := &chatMsg.Msg{
		From:           "luka1",
		Target:         "luka2",
		Content:        []byte("fuck"),
		MsgType:        1,
		MsgContentType: 1,
		EndConn:        false,
	}
	message, _ := proto.Marshal(m)
	log.Println(message)
	_ = pod.Socket.SendOneToSocket(message)
}
