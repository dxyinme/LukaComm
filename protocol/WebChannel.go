package protocol

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dxyinme/LukaComm/util/config"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
	"log"
)

type ChannelServer struct {

	StopListener context.CancelFunc

	listenerContext context.Context

	listener *kcp.Listener

	handlerFuncMap map[string] func(in []byte)[]byte
}

func (cs *ChannelServer) serveHandler(f func(in []byte) []byte, conn *KcpSocket) error {
	in ,err := conn.ReadOneFromSocket()
	if err != nil {
		log.Println(err)
	} else {
		errSend := conn.SendOneToSocket(f(in))
		if errSend != nil {
			return errSend
		}
		Ack, errAck := conn.ReadOneFromSocket()
		if errAck != nil {
			return errAck
		}
		if string(Ack) != ConfirmCode + conn.GetName() {
			return fmt.Errorf("error Ack , want %s , but %s", ConfirmCode + conn.GetName(), string(Ack))
		}
	}
	return nil
}

// 同样的handler会被后面来的顶替掉
func (cs *ChannelServer) HandlerFunc(handler string, f func(in []byte)[]byte) {
	cs.handlerFuncMap[handler] = f
}

func (cs *ChannelServer) ServeMultiConn() error {
	for {
		socketNow, errNow := cs.listener.AcceptKCP()
		if errNow != nil {
			continue
		}
		nowKcpSocket := &KcpSocket{conn: socketNow}
		// 设置超时时间为1s
		_ = nowKcpSocket.SetLiveTime(UsualKCPLiveTime)

		errSocket := nowKcpSocket.Initial()
		if errSocket != nil {
			continue
		}
		errCheck := nowKcpSocket.SendOneToSocket([]byte(ConfirmCode + nowKcpSocket.GetName()))
		if errCheck != nil {
			continue
		}
		if handlerNow,ok := cs.handlerFuncMap[nowKcpSocket.GetName()] ; ok {
			go func () {
				err := cs.serveHandler(handlerNow, nowKcpSocket)
				if err != nil {
					log.Println(err)
				}
			}()
		} else {
			log.Printf("has no such handler %s",nowKcpSocket.GetName())
		}
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
	cs.handlerFuncMap = make(map[string]func(in []byte)[]byte)
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


type ChannelClient struct {
	dial 	*Dialer
}

func (cc *ChannelClient) ClientCall(rAddr, password, saltWord, handler string, request []byte) ([]byte,error) {
	var (
		err error
		ret []byte
	)
	cc.dial = &Dialer{}
	err = cc.dial.Dial(rAddr,password,saltWord,handler,10,3)
	if err != nil {
		return nil,err
	}
	err = cc.dial.Socket.SendOneToSocket(request)
	if err != nil {
		return nil,err
	}
	ret,err = cc.dial.Socket.ReadOneFromSocket()
	if err != nil {
		return nil,err
	}
	err = cc.dial.Socket.SendOneToSocket([]byte(ConfirmCode + handler))
	if err != nil {
		return nil,err
	}
	err = cc.dial.Socket.conn.Close()
	if err != nil {
		return nil,err
	}
	return ret,nil
}