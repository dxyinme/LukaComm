package Client

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	timeout time.Duration
	conn 	*grpc.ClientConn
	client chatMsg.MsgCynicClient
}
