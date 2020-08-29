package main

import (
	"github.com/dxyinme/LukaComm/LukaChannel"
	"log"
)

func main(){
	pod := LukaChannel.Dialer{}
	err := pod.Dial("127.0.0.1:33456", "xxx","xxx2","luka1",10,3)
	if err != nil {
		log.Println(err)
	}
}
