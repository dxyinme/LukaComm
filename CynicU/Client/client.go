package Client

import (
	"context"
	"fmt"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/MD5"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	timeout time.Duration
	conn 	*grpc.ClientConn
	client  chatMsg.MsgCynicClient
	host 	string
}

func (c *Client) SendTo(msg *chatMsg.Msg) error {
	nowContext,cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	var (
		md5Msg string
		pbMsg []byte
		err error
		resp *chatMsg.Ack
	)
	resp,err = c.client.SendTo(nowContext, msg)
	if err != nil {
		return err
	}
	pbMsg,err = proto.Marshal(msg)
	md5Msg = MD5.CalcMD5(pbMsg)
	if resp.Md5 != md5Msg {
		return fmt.Errorf("error ack")
	}
	return nil
}

func (c *Client) Pull(req *chatMsg.PullReq) (*chatMsg.MsgPack, error){
	nowContext,cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	resp, err := c.client.Pull(nowContext, req)
	if err != nil {
		return nil,err
	}
	return resp, nil
}

func (c *Client) PullAll(req *chatMsg.PullReq) (*chatMsg.MsgPack, error) {
	nowContext,cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	resp, err := c.client.PullAll(nowContext, req)
	if err != nil {
		if err != fmt.Errorf("NoMessage") {
			// reconnect in RPC FAIL time
			err = c.reconnect()
			if err != nil {
				return nil, err
			}
			resp, err = c.client.PullAll(nowContext, req)
			return resp, nil
		} else {
			return nil, err
		}
	}
	return resp,nil
}

func (c *Client) Initial(host string, timeout time.Duration) error {
	var (
		err error
	)
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.host = host
	c.conn = conn
	c.client = chatMsg.NewMsgCynicClient(c.conn)
	c.timeout = timeout
	return nil
}

func (c *Client) reconnect() error {
	conn,err := grpc.Dial(c.host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.Close()
	c.conn = conn
	c.client = chatMsg.NewMsgCynicClient(c.conn)
	return nil
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}