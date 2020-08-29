package main

import (
	"github.com/dxyinme/LukaComm/LukaChannel"
	"github.com/dxyinme/LukaComm/util/config"
	"log"
)

func main() {
	config.SetSelfIP("127.0.0.1:33456")
	pod := LukaChannel.ChannelServer{}
	err := pod.Initial("xxx","xxx2",10,3)
	if err != nil {
		log.Println(err)
	}
	if errServe := pod.ServeMultiConn() ; errServe != nil {
		log.Println(errServe)
	}
}