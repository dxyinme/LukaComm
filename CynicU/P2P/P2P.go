package P2P

import (
	"errors"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/MD5"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"time"
)

const (
	PacketSize = 1024
	AckSize = 32
)

var (
	SendTimeoutErr = errors.New("send message timeout")
	NotCorrectAckErr = errors.New("error ack")
)

type P2P struct {

	ServeAddr *net.UDPAddr

	RecvMsg chan *chatMsg.Msg

	conn *net.UDPConn

	sendConn net.Conn

	isListening bool
}

func (p *P2P) Listen(addr string) (err error) {
	p.ServeAddr,err = net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	p.conn, err = net.ListenUDP("udp", p.ServeAddr)
	if err != nil {
		return err
	}

	p.isListening = true
	go p.run()
	return nil
}

func (p *P2P) run() {
	var (
		msg []byte
		rAddr *net.UDPAddr
		err error
	)
	for {
		msg = make([]byte, PacketSize)
		_,rAddr,err = p.conn.ReadFromUDP(msg)
		msg = clarify(msg)
		if err != nil {
			log.Println(err)
		}
		var nowMsg chatMsg.Msg
		err = proto.Unmarshal(msg, &nowMsg)
		if err != nil {
			log.Println(err)
		}
		_,err = p.conn.WriteTo([]byte(MD5.CalcMD5(msg)), rAddr)
		p.RecvMsg <- &nowMsg
	}
}

func (p *P2P) Send(msg *chatMsg.Msg, addr string) (err error) {
	var finished = make(chan error, 1)
	defer close(finished)
	p.sendConn, err = net.Dial("udp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		var se error
		b, se := proto.Marshal(msg)
		if se != nil {
			finished <- se
			return
		}
		_,se = p.sendConn.Write(b)
		if se != nil {
			finished <- se
			return
		}
		ack := make([]byte, AckSize)
		_, se = p.sendConn.Read(ack)
		if se != nil {
			finished <- se
			return
		}
		if string(ack) != MD5.CalcMD5(b) {
			finished <- NotCorrectAckErr
			return
		}
		finished<-nil
	}()
	select {
	case err = <-finished:
		break
	case <-time.After(3 * time.Second):
		err = SendTimeoutErr
		break
	}
	return err
}

func NewP2P() *P2P {
	return &P2P{
		ServeAddr:   nil,
		RecvMsg:     make(chan *chatMsg.Msg, 10),
		conn:        nil,
		isListening: false,
	}
}

func clarify(b []byte) []byte {
	n := len(b)
	for ; n > 0 && b[n - 1] == 0; n -- {

	}
	return b[:n]
}