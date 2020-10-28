package Server

import (
	"context"
	"github.com/dxyinme/LukaComm/Assigneer"
	"github.com/dxyinme/LukaComm/util/CoHash"
	"github.com/golang/glog"
)

type Server struct {
	assignToStruct CoHash.AssignToStruct
	// KV :
	// key [ keeperID ] , value [ host ]
	hosts map[uint32]string
}

func (s *Server) SyncLocation(context.Context, *Assigneer.SyncLocationReq) (*Assigneer.SyncLocationRsp, error) {
	var ret []uint32
	var retHost []string
	for i := 0 ; i < len(s.assignToStruct.KeeperIDs); i ++ {
		ret = append(ret, uint32(s.assignToStruct.KeeperIDs[i]))
	}
	for i := 0 ; i < len(ret); i ++ {
		retHost = append(retHost, s.hosts[ret[i]])
	}
	return &Assigneer.SyncLocationRsp{
		KeeperIDs: 	ret,
		Hosts:		retHost,
	},nil
}

func (s *Server) RemoveKeeper(ctx context.Context, in *Assigneer.RemoveKeeperReq) (*Assigneer.AssignAck, error) {
	err := s.assignToStruct.RemoveKeeper(in.KeeperID)
	if err != nil {
		glog.Info(err)
	}
	delete(s.hosts, in.KeeperID)
	return &Assigneer.AssignAck{
		AckMessage: "",
	},nil
}

func (s *Server) AddKeeper(ctx context.Context,in *Assigneer.AddKeeperReq) (*Assigneer.AssignAck, error) {
	s.assignToStruct.AppendKeeper(in.KeeperID)
	s.hosts[in.KeeperID] = in.Host
	return &Assigneer.AssignAck{
		AckMessage: "",
	},nil
}

