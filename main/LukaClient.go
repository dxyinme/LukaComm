package main

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/protocol"
	"github.com/golang/protobuf/proto"
	"log"
)

func main(){
	m := &chatMsg.Msg{
		From:           "luka1",
		Target:         "luka2",
		Content:        []byte("fuck"),
		MsgType:        1,
		MsgContentType: 1,
	}
	message, _ := proto.Marshal(m)
	log.Println(message)
	Client := protocol.ChannelClient{}
	result, err := Client.ClientCall("127.0.0.1:33456", "xxx","xxx2","luka1",message)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(result))
}
