package SendMsg

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/MD5"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

const (
	PacketSize = 512
	ChannelSize = 10
)

type Server struct {
	conn *net.UDPConn
	localAddr *net.UDPAddr
	msgCh chan *chatMsg.Msg
}

func (s *Server) GetMsg() *chatMsg.Msg {
	select {
	case ret := <- s.msgCh:
		return ret
	}
}

func (s *Server) Handle(b []byte, rAddr *net.UDPAddr) {
	log.Println(rAddr.String())
	msgNow := &chatMsg.Msg{}
	md5 := []byte(MD5.CalcMD5(b))
	err := proto.Unmarshal(b, msgNow)
	if err != nil {
		log.Println(msgNow, err)
		return
	}
	s.msgCh <- msgNow

	_, err = s.conn.WriteToUDP(md5, rAddr)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) Listen(addr string) (err error) {
	s.localAddr, err = net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	s.conn, err = net.ListenUDP("udp", s.localAddr)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("listen prepare finished")
	for {
		var (
			rAddr *net.UDPAddr
			n int
		)
		b := make([]byte, PacketSize)
		n, rAddr, err = s.conn.ReadFromUDP(b)
		log.Println(len(b))
		if err != nil {
			log.Println(err)
			break
		}
		if n <= 0 {
			continue
		} else {
			go s.Handle(b[:n], rAddr)
		}
	}
	return
}


func NewServer() *Server {
	return &Server{
		conn:      nil,
		localAddr: nil,
		msgCh:     make(chan *chatMsg.Msg, ChannelSize),
	}
}