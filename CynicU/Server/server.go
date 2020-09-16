package Server

import (
	"context"
	"github.com/dxyinme/LukaComm/chatMsg"
	"google.golang.org/grpc"
	"log"
	"net"
)

type WorkerPool interface {
	SendTo(msg *chatMsg.Msg)
	Pull(targetIs string) (*chatMsg.MsgPack,error)
}

type Server struct {
	Lis net.Listener
	name string
	w WorkerPool
}

func (s *Server) SendTo(ctx context.Context, in *chatMsg.Msg) (*chatMsg.Ack, error) {
	defer func(){
		go s.w.SendTo(in)
	}()
	return &chatMsg.Ack{From: in.From},nil
}

func (s *Server) Pull(ctx context.Context, in *chatMsg.Ack) (*chatMsg.MsgPack, error) {
	return s.w.Pull(in.From)
}

func (s *Server) NewCynicUServer(addr string, name string) *grpc.Server {
	var (
		err error
		retServer *grpc.Server
	)
	// to avoid Running Error
	s.w = &UnImplWorkerPool{}

	s.Lis, err = net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("%s failed to listen: %v", name, err)
	}
	s.name = name
	retServer = grpc.NewServer()
	chatMsg.RegisterMsgCynicServer(retServer, s)
	return retServer
}

func (s *Server) BindWorkerPool(pool WorkerPool) {
	s.w = pool
}


type UnImplWorkerPool struct {
}
func (uiw *UnImplWorkerPool) SendTo(msg *chatMsg.Msg) {
	log.Println("No Impl for WorkerPool - func SendTo")
}

func (uiw *UnImplWorkerPool) Pull(targetIs string) (*chatMsg.MsgPack,error) {
	log.Println("No Impl for WorkerPool - func Pull")
	return nil,nil
}