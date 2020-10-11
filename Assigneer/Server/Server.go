package Server

import (
	"context"
	"github.com/dxyinme/LukaComm/Assigneer"
	"github.com/dxyinme/LukaComm/util/CoHash"
	"github.com/golang/glog"
)

type Server struct {
	assignToStruct CoHash.AssignToStruct
}

func (s *Server) SyncLocation(context.Context, *Assigneer.SyncLocationReq) (*Assigneer.SyncLocationRsp, error) {
	var ret []uint32
	for i := 0 ; i < len(s.assignToStruct.KeeperIDs); i ++ {
		ret = append(ret, uint32(s.assignToStruct.KeeperIDs[i]))
	}
	return &Assigneer.SyncLocationRsp{
		KeeperIDs: ret,
	},nil
}

func (s *Server) RemoveKeeper(ctx context.Context, in *Assigneer.RemoveKeeperReq) (*Assigneer.AssignAck, error) {
	err := s.assignToStruct.RemoveKeeper(in.KeeperID)
	if err != nil {
		glog.Info(err)
	}
	return &Assigneer.AssignAck{
		AckMessage: "",
	},nil
}

func (s *Server) AddKeeper(ctx context.Context,in *Assigneer.AddKeeperReq) (*Assigneer.AssignAck, error) {
	s.assignToStruct.AppendKeeper(in.KeeperID)
	return &Assigneer.AssignAck{
		AckMessage: "",
	},nil
}

