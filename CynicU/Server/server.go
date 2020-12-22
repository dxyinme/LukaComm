package Server

import (
	"context"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/MD5"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

type WorkerPool interface {
	// send message.
	SendTo(msg *chatMsg.Msg)
	// daily use pull user's message.
	Pull(targetIs string) (*chatMsg.MsgPack, error)
	// do not use PullAll but cluster is updating.
	PullAll(targetIs string) (*chatMsg.MsgPack, error)
	// SyncLocation
	SyncLocationNotify()
	// leave from group
	LeaveGroup(req *chatMsg.GroupReq) error
	// join into group
	JoinGroup(req *chatMsg.GroupReq) error
	// create group , the uid is group master
	CreateGroup(req *chatMsg.GroupReq) error
	// delete group , the uid is group master
	DeleteGroup(req *chatMsg.GroupReq) error
	// checkAlive , get the information of keeper.
	CheckAlive(req *chatMsg.KeepAlive) *chatMsg.KeepAlive
	// usage Call
	UseCall(channel *chatMsg.UseChannel) (*chatMsg.UseChannel, error)
}

type Server struct {
	Lis  net.Listener
	name string
	w    WorkerPool
}

func (s *Server) UseCall(ctx context.Context,in *chatMsg.UseChannel) (ret *chatMsg.UseChannel,err error) {
	ret, err = s.w.UseCall(in)
	return
}

func (s *Server) CreateGroup(ctx context.Context, in *chatMsg.GroupReq) (*chatMsg.GroupRsp, error) {
	err := s.w.CreateGroup(in)
	if err != nil {
		return &chatMsg.GroupRsp{
			Msg: err.Error(),
		}, nil
	} else {
		return &chatMsg.GroupRsp{
			Msg: "",
		}, nil
	}
}

func (s *Server) DeleteGroup(ctx context.Context, in *chatMsg.GroupReq) (*chatMsg.GroupRsp, error) {
	err := s.w.DeleteGroup(in)
	if err != nil {
		return &chatMsg.GroupRsp{
			Msg: err.Error(),
		}, nil
	} else {
		return &chatMsg.GroupRsp{
			Msg: "",
		}, nil
	}
}

func (s *Server) JoinGroup(ctx context.Context, in *chatMsg.GroupReq) (*chatMsg.GroupRsp, error) {
	err := s.w.JoinGroup(in)
	if err != nil {
		return &chatMsg.GroupRsp{
			Msg: err.Error(),
		}, nil
	} else {
		return &chatMsg.GroupRsp{
			Msg: "",
		}, nil
	}
}

func (s *Server) LeaveGroup(ctx context.Context, in *chatMsg.GroupReq) (*chatMsg.GroupRsp, error) {
	err := s.w.LeaveGroup(in)
	if err != nil {
		return &chatMsg.GroupRsp{
			Msg: err.Error(),
		}, nil
	} else {
		return &chatMsg.GroupRsp{
			Msg: "",
		}, nil
	}
}

func (s *Server) SyncLocationNotify(context.Context, *chatMsg.KeepAlive) (*chatMsg.KeepAlive, error) {
	s.w.SyncLocationNotify()
	return &chatMsg.KeepAlive{
		CheckAlive: "",
	}, nil
}

func (s *Server) PullAll(ctx context.Context, in *chatMsg.PullReq) (*chatMsg.MsgPack, error) {
	return s.w.PullAll(in.From)
}

func (s *Server) SendTo(ctx context.Context, in *chatMsg.Msg) (*chatMsg.Ack, error) {
	defer func() {
		go s.w.SendTo(in)
	}()
	pbIn, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	return &chatMsg.Ack{Md5: MD5.CalcMD5(pbIn)}, nil
}

func (s *Server) Pull(ctx context.Context, in *chatMsg.PullReq) (*chatMsg.MsgPack, error) {
	return s.w.Pull(in.From)
}

func (s *Server) CheckAlive(ctx context.Context, in *chatMsg.KeepAlive) (*chatMsg.KeepAlive, error) {
	ret := s.w.CheckAlive(in)
	return ret, nil
}

func (s *Server) NewCynicUServer(addr string, name string) *grpc.Server {
	var (
		err       error
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

func (uiw *UnImplWorkerPool) UseCall(channel *chatMsg.UseChannel) (*chatMsg.UseChannel, error) {
	log.Println("No Impl for WorkerPool - func UseCall")
	return nil,nil
}

func (uiw *UnImplWorkerPool) CheckAlive(req *chatMsg.KeepAlive) *chatMsg.KeepAlive {
	log.Println("No Impl for WorkerPool - func CheckAlive")
	return nil
}

func (uiw *UnImplWorkerPool) DeleteGroup(req *chatMsg.GroupReq) error {
	log.Println("No Impl for WorkerPool - func DeleteGroup")
	return nil
}

func (uiw *UnImplWorkerPool) CreateGroup(req *chatMsg.GroupReq) error {
	log.Println("No Impl for WorkerPool - func CreateGroup")
	return nil
}

func (uiw *UnImplWorkerPool) LeaveGroup(req *chatMsg.GroupReq) error {
	log.Println("No Impl for WorkerPool - func LeaveGroup")
	return nil
}

func (uiw *UnImplWorkerPool) JoinGroup(req *chatMsg.GroupReq) error {
	log.Println("No Impl for WorkerPool - func JoinGroup")
	return nil
}

func (uiw *UnImplWorkerPool) SendTo(msg *chatMsg.Msg) {
	log.Println("No Impl for WorkerPool - func SendTo")
}

func (uiw *UnImplWorkerPool) Pull(targetIs string) (*chatMsg.MsgPack, error) {
	log.Println("No Impl for WorkerPool - func Pull")
	return &chatMsg.MsgPack{}, nil
}

func (uiw *UnImplWorkerPool) PullAll(targetIs string) (*chatMsg.MsgPack, error) {
	log.Println("No Impl for WorkerPool - func PullAll")
	return &chatMsg.MsgPack{}, nil
}

func (uiw *UnImplWorkerPool) SyncLocationNotify() {
	log.Println("No Impl for WorkerPool - func SyncLocationNotify")
}
