package userClient

import (
	"context"
	"github.com/dxyinme/LukaComm/Assigneer"
	"google.golang.org/grpc"
)


// only SwitchKeeper will be used in client.
func SwitchKeeper(req *Assigneer.SwitchKeeperReq) (rsp *Assigneer.SwitchKeeperRsp, err error) {
	conn, err := grpc.Dial("client", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := Assigneer.NewAssigneerClient(conn)
	rsp, err = c.SwitchKeeper(context.Background(), req)
	return
}