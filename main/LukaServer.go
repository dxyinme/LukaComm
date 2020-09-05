package main

import (
	"github.com/dxyinme/LukaComm/LukaChannel"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/config"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	config.SetSelfIP("127.0.0.1:33456")
	pod := LukaChannel.ChannelServer{}
	err := pod.Initial("xxx","xxx2",10,3)
	if err != nil {
		log.Println(err)
	}
	go func() {
		var o *LukaChannel.KcpSocket
		for {
			o = pod.GetSocket("luka1")
			if o != nil {
				break
			}
		}
		f,_ := o.ReadOneFromSocket()
		log.Println(f)
		var m *chatMsg.Msg
		m = &chatMsg.Msg{}
		err := proto.Unmarshal(f, m)
		if err != nil {
			log.Println(err)
		}
		log.Println(m)
	}()
	if errServe := pod.ServeMultiConn() ; errServe != nil {
		log.Println(errServe)
	}
}