package main

import (
	CynicUServer "github.com/dxyinme/LukaComm/CynicU/Server"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"os"
	"sync"
)



var (
	mu sync.Mutex
	recvCnt int64 = 0
)

type wk struct {
}

func (w *wk) CheckAlive(req *chatMsg.KeepAlive) *chatMsg.KeepAlive {
	panic("implement me")
}

func (w *wk) UseCall(channel *chatMsg.UseChannel) (*chatMsg.UseChannel, error) {
	panic("implement me")
}

func (w *wk) SyncLocationNotify() {
	panic("implement me")
}

func (w *wk) LeaveGroup(req *chatMsg.GroupReq) error {
	panic("implement me")
}

func (w *wk) JoinGroup(req *chatMsg.GroupReq) error {
	panic("implement me")
}

func (w *wk) CreateGroup(req *chatMsg.GroupReq) error {
	panic("implement me")
}

func (w *wk) DeleteGroup(req *chatMsg.GroupReq) error {
	panic("implement me")
}

func (w *wk) SendTo(msg *chatMsg.Msg) {
	log.Println(msg)
	mu.Lock()
	defer mu.Unlock()
	recvCnt ++
	log.Println(recvCnt)
}

func (w *wk) Pull(targetIs string) (*chatMsg.MsgPack, error) {
	//log.Println(targetIs)
	return &chatMsg.MsgPack{}, nil
}

func (w *wk) PullAll(targetIs string) (*chatMsg.MsgPack, error) {
	return &chatMsg.MsgPack{}, nil
}

func main() {
	log.Printf("PID=%d\n", os.Getpid())
	s := &CynicUServer.Server{}
	gServe := s.NewCynicUServer(":8080", "Luka")
	s.BindWorkerPool(&wk{})
	if err := gServe.Serve(s.Lis); err != nil {
		log.Fatal(err)
	}
}
