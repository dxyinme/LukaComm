package Client

import (
	"context"
	"errors"
	"fmt"
	"github.com/dxyinme/LukaComm/Auth"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/Const"
	"github.com/dxyinme/LukaComm/util/MD5"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"time"
)

type Client struct {
	timeout time.Duration
	conn    *grpc.ClientConn
	client  chatMsg.MsgCynicClient
	host    string
}

func (c *Client) SendTo(msg *chatMsg.Msg) error {
	nowContext, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	var (
		md5Msg string
		pbMsg  []byte
		err    error
		resp   *chatMsg.Ack
	)
	msg.SendTime = time.Now().String()
	resp, err = c.client.SendTo(nowContext, msg)
	if err != nil {
		return err
	}
	pbMsg, err = proto.Marshal(msg)
	if err != nil {
		return err
	}
	md5Msg = MD5.CalcMD5(pbMsg)
	if resp.Md5 != md5Msg {
		return fmt.Errorf("error ack")
	}
	return nil
}

func (c *Client) SendKeyAgreement(from, target string, keyAfterRsa []byte) error {
	keyAgreementMsg := &chatMsg.Msg{
		From:           from,
		Target:         target,
		Content:        keyAfterRsa,
		MsgType:        chatMsg.MsgType_Single,
		MsgContentType: chatMsg.MsgContentType_KeyAgreement,
		SendTime:       time.Now().String(),
	}
	return c.SendTo(keyAgreementMsg)
}

func (c *Client) GetAuthPubKey(host, uid string) ([]byte,error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	nowContext, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	authClient := Auth.NewAuthClient(conn)
	rsp, err := authClient.GetAuthPubKey(nowContext, &Auth.GetAuthPubKeyReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	if rsp == nil {
		return nil, errors.New("no response")
	}
	return rsp.AuthRsaPubKey, nil
}

func (c *Client) SetAuthPubKey(host, uid string, pubKey []byte) error {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	nowContext, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	authClient := Auth.NewAuthClient(conn)
	rsp, err := authClient.SetAuthPubKey(nowContext, &Auth.SetAuthPubKeyReq{
		Uid: uid,
		AuthRsaPubKey: pubKey,
	})
	if err != nil {
		return err
	}
	if rsp == nil {
		return errors.New("no response")
	}
	return errors.New(rsp.ErrorMsg)
}

func (c *Client) SendToInSecret(msg *chatMsg.Msg) error {
	msg.SecretLevel = 1
	return c.SendTo(msg)
}

func (c *Client) Pull(req *chatMsg.PullReq) (*chatMsg.MsgPack, error) {
	nowContext, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	resp, err := c.client.Pull(nowContext, req)
	return resp, err
}

func (c *Client) PullAll(req *chatMsg.PullReq) (*chatMsg.MsgPack, error) {
	nowContext, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	resp, err := c.client.PullAll(nowContext, req)
	return resp, err
}

func (c *Client) CheckAlive() (ret *chatMsg.KeepAlive,err error) {
	ret, err = c.client.CheckAlive(
		context.Background(),
		&chatMsg.KeepAlive{
			CheckAlive: time.Now().String(),
		})
	return
}

// opNum operator is in LukaComm/util/Const
func (c *Client) GroupOp(opNum string, req *chatMsg.GroupReq) error {
	var (
		err error
		rsp *chatMsg.GroupRsp
	)
	switch opNum {
	case Const.CreateGroup:
		rsp, err = c.client.CreateGroup(context.Background(), req)
		break
	case Const.DeleteGroup:
		rsp, err = c.client.DeleteGroup(context.Background(), req)
		break
	case Const.JoinGroup:
		rsp, err = c.client.JoinGroup(context.Background(), req)
		break
	case Const.LeaveGroup:
		rsp, err = c.client.LeaveGroup(context.Background(), req)
		break
	}
	if err != nil {
		return err
	}
	if rsp.Msg != "" {
		return fmt.Errorf("error rsp : %s", rsp.Msg)
	}
	return nil
}

func (c *Client) SyncLocationNotify() error {
	_, err := c.client.SyncLocationNotify(context.Background(), &chatMsg.KeepAlive{
		CheckAlive: "",
	})
	return err
}

func (c *Client) UseCall(in *chatMsg.UseChannel) (ret *chatMsg.UseChannel, err error) {
	ret, err = c.client.UseCall(context.Background(), in)
	return
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
	conn, err := grpc.Dial(c.host, grpc.WithInsecure())
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
		_ = c.conn.Close()
	}
}
