package Client

import (
	"context"
	"github.com/dxyinme/LukaComm/chatMsg"
	"google.golang.org/grpc"
	"time"
)

type StreamClient struct {
	host 			string
	conn 			*grpc.ClientConn
	client 			chatMsg.MsgCynicClient
	streamClient 	chatMsg.MsgCynic_MsgConnClient
}

func (sc *StreamClient) Initial(uid string, host string, timeout time.Duration) error {
	var (
		err error
	)
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	sc.host = host
	sc.conn = conn
	sc.client = chatMsg.NewMsgCynicClient(sc.conn)
	sc.streamClient,err = sc.client.MsgConn(context.Background())
	if err != nil {
		return err
	}
	err = sc.streamClient.Send(&chatMsg.Msg{
		From:           uid,
	})
	if err != nil {
		return err
	}
	return nil
}

func (sc *StreamClient) Send(msg *chatMsg.Msg) error {
	return sc.streamClient.Send(msg)
}

func (sc *StreamClient) Recv() (*chatMsg.MsgPack, error) {
	return sc.streamClient.Recv()
}

func (sc *StreamClient) CheckAlive() bool {
	_, err := sc.client.CheckAlive(context.Background(), &chatMsg.KeepAlive{
		CheckAlive: "",
	})
	return err != nil
}