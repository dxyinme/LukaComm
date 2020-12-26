package SendMsg

import (
	"errors"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/MD5"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"time"
)

type Client struct {
	conn net.Conn

	addr string
}

func (c *Client) SendTo(msg *chatMsg.Msg) (err error) {
	var (
		md5 string
		n int
		b []byte
		ack = make([]byte, 32)
		errCh chan error
		finCh chan error
	)
	finCh = make(chan error, 1)
	errCh = make(chan error, 1)
	f := func() {
		var errF error
		c.conn, errF = net.Dial("udp", c.addr)
		if err != nil {
			log.Println(err)
			return
		}
		defer c.conn.Close()
		b, errF = proto.Marshal(msg)
		if errF != nil {
			errCh <- errF
			return
		}
		if len(b) <= PacketSize {
			md5 = MD5.CalcMD5(b)
			_, errF = c.conn.Write(b)
			if errF != nil {
				errCh <- errF
				return
			}
		} else {
			errCh <- errors.New("too long msg")
		}
		n, err = c.conn.Read(ack)
		if n != 32 {
			errCh <- errors.New("ack length error , not equal to 32")
		}

		if string(ack) != md5 {
			errCh <- errors.New("ack content error")
		}
		finCh <- nil
	}
	go f()
	select {
	case err = <- errCh:
		log.Println(err)
		break
	case <- time.After(1 * time.Second):
		err = errors.New("time out")
		break
	case <- finCh:
		break
	}
	return
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}