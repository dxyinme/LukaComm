package protocol

import (
	"crypto/sha1"
	"fmt"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
)

type Dialer struct {
	name string
	Socket *KcpSocket
}

// notice : 这里的name值的是handler名，不是连接的名字。
func (dialer *Dialer) Dial(rAddr, password, saltWord, name string, dataShards, parityShards int) error {
	var (
		ackMsg []byte
		err error
		nowConn *kcp.UDPSession
	)
	key := pbkdf2.Key([]byte(password), []byte(saltWord), 1024, 32, sha1.New)
	block, err := kcp.NewAESBlockCrypt(key)
	if err != nil {
		return err
	}
	dialer.name = name
	nowConn, err = kcp.DialWithOptions(rAddr, block, dataShards, parityShards)
	if err != nil {
		return err
	}
	dialer.Socket = &KcpSocket{conn: nowConn}
	err = dialer.Socket.SetLiveTime(UsualKCPLiveTime)
	if err != nil {
		return err
	}
	err = dialer.Socket.SendOneToSocket([]byte(dialer.name))
	if err != nil {
		return err
	}
	ackMsg, err = dialer.Socket.ReadOneFromSocket()
	if err != nil {
		return err
	}
	if string(ackMsg) != ConfirmCode + dialer.name {
		return fmt.Errorf("no ack for socket %s", dialer.name)
	}
	return nil
}