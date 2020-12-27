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

var (
	TimeoutErr = errors.New("time out")
	TooLongMsgErr = errors.New("too long msg")
	AckLengthErr = errors.New("ack length error , not equal to 32")
	AckContentErr = errors.New("ack content error")
)

type Client struct {
	conn net.Conn

	TimeoutLimit time.Duration

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
			errCh <- TooLongMsgErr
		}
		n, err = c.conn.Read(ack)
		if n != 32 {
			errCh <- AckLengthErr
		}

		if string(ack) != md5 {
			errCh <- AckContentErr
		}
		close(finCh)
	}
	go f()
	select {
	case err = <- errCh:
		log.Println(err)
		break
	case <- time.After(c.TimeoutLimit):
		err = TimeoutErr
		break
	case <- finCh:
		break
	}
	return
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
		TimeoutLimit: 1 * time.Second,
	}
}