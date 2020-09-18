package Client

import (
	"context"
	"fmt"
	"github.com/dxyinme/LukaComm/chatMsg"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	timeout time.Duration
	conn 	*grpc.ClientConn
	client  chatMsg.MsgCynicClient
}

func (c *Client) SendTo(msg *chatMsg.Msg) error {
	nowContext,cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	resp,err := c.client.SendTo(nowContext, msg)
	if err != nil {
		return err
	}
	if resp.From != msg.From {
		return fmt.Errorf("error Ack , want %s, but %s", msg.From, resp.From)
	}
	return nil
}

func (c *Client) Pull(ack *chatMsg.Ack) (*chatMsg.MsgPack, error){
	nowContext,cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	resp, err := c.client.Pull(nowContext, ack)
	if err != nil {
		return nil,err
	}
	return resp, nil
}

func (c *Client) PullAll(ack *chatMsg.Ack) (*chatMsg.MsgPack, error) {

	return nil,nil
}

func (c *Client) Initial(host string, timeout time.Duration) error {
	var (
		err error
	)
	conn, err := grpc.Dial(host,grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.conn = conn
	c.client = chatMsg.NewMsgCynicClient(c.conn)
	c.timeout = timeout
	return nil
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}