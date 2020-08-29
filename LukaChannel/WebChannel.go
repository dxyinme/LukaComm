package LukaChannel

import (
	"context"
	"crypto/sha1"
	"github.com/dxyinme/LukaComm/util/config"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
	"log"
)

type ChannelServer struct {

	StopListener context.CancelFunc

	listenerContext context.Context

	listener *kcp.Listener

	connPool map[string]*KcpSocket
}

func (cs *ChannelServer) ServeMultiConn() error {
	for {
		socketNow, errNow := cs.listener.AcceptKCP()
		if errNow != nil {
			continue
		}
		nowKcpSocket := &KcpSocket{conn: socketNow}
		errSocket := nowKcpSocket.Initial()
		if errSocket != nil {
			continue
		}
		log.Println(nowKcpSocket.GetName())
		errCheck := nowKcpSocket.SendOneToSocket([]byte(ConfirmCode + nowKcpSocket.GetName()))
		if errCheck != nil {
			continue
		}
		cs.connPool[nowKcpSocket.GetName()] = nowKcpSocket
		select {
		case <-cs.listenerContext.Done():
			return nil
		}
	}
}

func (cs *ChannelServer) Initial(password, saltWord string, dataShards, parityShards int) error{
	var (
		address string
		err error
	)
	cs.connPool = make(map[string]*KcpSocket)
	key := pbkdf2.Key([]byte(password), []byte(saltWord), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	address = config.GetSelfIP()
	cs.listener, err = kcp.ListenWithOptions(address, block, dataShards, parityShards)
	if err != nil {
		return err
	}
	cs.listenerContext, cs.StopListener = context.WithCancel(context.Background())
	return nil
}