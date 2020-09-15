package Server

import (
	"context"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
)

type Server struct {
	
}

func (s *Server) SendTo(ctx context.Context, in *chatMsg.Msg) (*chatMsg.Ack, error) {
	log.Println(in)
	return &chatMsg.Ack{From: in.From},nil
}

func (s *Server) Pull(ctx context.Context, in *chatMsg.Ack) (*chatMsg.MsgPack, error) {
	return nil,nil
}

