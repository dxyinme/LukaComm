package Server

import (
	"context"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/MD5"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type WorkerPool interface {
	// send message.
	SendTo(msg *chatMsg.Msg)
	// daily use pull user's message.
	Pull(targetIs string) (*chatMsg.MsgPack,error)
	// do not use PullAll but cluster is updating.
	PullAll(targetIs string) (*chatMsg.MsgPack,error)
}

type Server struct {
	Lis net.Listener
	name string
	w WorkerPool
}

func (s *Server) PullAll(ctx context.Context, in *chatMsg.PullReq) (*chatMsg.MsgPack, error) {
	return s.w.PullAll(in.From)
}

func (s *Server) SendTo(ctx context.Context, in *chatMsg.Msg) (*chatMsg.Ack, error) {
	defer func(){
		go s.w.SendTo(in)
	}()
	pbIn, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	return &chatMsg.Ack{Md5: MD5.CalcMD5(pbIn)},nil
}

func (s *Server) Pull(ctx context.Context, in *chatMsg.PullReq) (*chatMsg.MsgPack, error) {
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
	return &chatMsg.MsgPack{},nil
}

func (uiw *UnImplWorkerPool) PullAll(targetIs string) (*chatMsg.MsgPack,error) {
	log.Println("No Impl for WorkerPool - func PullAll")
	return &chatMsg.MsgPack{},nil
}