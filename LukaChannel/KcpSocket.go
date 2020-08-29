package LukaChannel

import (
	"github.com/xtaci/kcp-go/v5"
	"log"
)

type KcpSocket struct {
	name string
	conn *kcp.UDPSession
}

func (kcpSocket *KcpSocket) ReadOneFromSocket() ([]byte,error) {
	var (
		AllBuf []byte
		buf = make([]byte, SocketLength)
	)
	for {
		n, err := kcpSocket.conn.Read(buf)
		if err != nil {
			// 对于失败的 socket 不做提醒操作
			return nil,err
		}
		if n <= 1 {
			continue
		}
		AllBuf = append(AllBuf, buf[1:n]...)
		if buf[0] == byte(0) {
			break
		}
	}
	return AllBuf,nil
}

func (kcpSocket *KcpSocket) SendOneToSocket(msg []byte) error {
	var (
		limLen = len(msg)
		err error
	)
	log.Println(msg)
	for i := 0 ; i < limLen ; i += SocketLength - 1 {
		nowSend := []byte("")
		if i + SocketLength - 1 >= limLen {
			nowSend = append(nowSend , byte(0))
			nowSend = append(nowSend , msg[i:] ...)
		} else {
			nowSend = append(nowSend , byte(1))
			nowSend = append(nowSend , msg[i:i + SocketLength - 1] ...)
		}
		_, err = kcpSocket.conn.Write(nowSend)
		if err != nil {
			return err
		}
	}
	return nil
}

func (kcpSocket *KcpSocket) Initial() error {
	nameByte, err := kcpSocket.ReadOneFromSocket()
	if err != nil {
		return err
	}
	kcpSocket.name = string(nameByte)
	return nil
}

func (kcpSocket *KcpSocket) GetName() string {
	return kcpSocket.name
}