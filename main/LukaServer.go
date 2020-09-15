package main

import (
	"github.com/dxyinme/LukaComm/protocol"
	"github.com/dxyinme/LukaComm/util/config"
	"log"
)

func luka1(in []byte)[]byte {
	log.Println(string(in))
	return []byte("noooo")
}

func main() {
	config.SetSelfIP("127.0.0.1:33456")
	pod := protocol.ChannelServer{}
	err := pod.Initial("xxx","xxx2",10,3)
	if err != nil {
		log.Println(err)
	}
	pod.HandlerFunc("luka1",luka1)
	if errServe := pod.ServeMultiConn() ; errServe != nil {
		log.Println(errServe)
	}
}