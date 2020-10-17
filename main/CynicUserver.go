package main

import (
	CynicUServer "github.com/dxyinme/LukaComm/CynicU/Server"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
)

type wk struct {

}

func (w *wk) SendTo(msg *chatMsg.Msg) {
	log.Println(msg)
}

func (w *wk) Pull(targetIs string) (*chatMsg.MsgPack, error) {
	//log.Println(targetIs)
	return &chatMsg.MsgPack{}, nil
}

func (w *wk) PullAll(targetIs string) (*chatMsg.MsgPack, error) {
	return &chatMsg.MsgPack{}, nil
}

func main() {
	s := &CynicUServer.Server{}
	gServe := s.NewCynicUServer(":8080", "Luka")
	s.BindWorkerPool(&wk{})
	if err := gServe.Serve(s.Lis); err != nil {
		log.Fatal(err)
	}
}
