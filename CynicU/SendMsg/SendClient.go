package SendMsg

import (
	"errors"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/MD5"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"sync"
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

	mu_ sync.Mutex

	nowPacketSize int

	lossTime int

	allTime int

	retry int

	modifyInterval int
}

func (c *Client) modifyNowPacketSize() {
	if c.lossTime == 0 {
		c.nowPacketSize *= 2
		if c.nowPacketSize > ServerPacketSize {
			c.nowPacketSize = ServerPacketSize
		}
	} else if c.lossTime < 5 {
		// nothing to do
	} else {
		c.nowPacketSize /= 2
	}
	c.lossTime = 0
	c.allTime = 0
}

func (c *Client) SendTo(msg *chatMsg.Msg) (err error) {
	c.mu_.Lock()
	defer c.mu_.Unlock()
	c.allTime ++
	var (
		md5 string
		n int
		b []byte
		errCh chan error
		finCh chan error
	)
	errCh = make(chan error, c.retry)
	finCh = make(chan error, c.retry)
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
		if len(b) <= c.nowPacketSize {
			for retryTime := 0 ; retryTime < c.retry ; retryTime ++ {
				md5 = MD5.CalcMD5(b)
				_, errF = c.conn.Write(b)
				if errF != nil {
					c.lossTime ++
				} else {
					break
				}
			}
		} else {
			errCh <- TooLongMsgErr
		}
		if errF != nil {
			errCh <- errF
		}
		var ack = make([]byte, 32)
		n, err = c.conn.Read(ack)
		if n != 32 {
			errCh <- AckLengthErr
			return
		}
		if string(ack) != md5 {
			errCh <- AckContentErr
			return
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
	if c.allTime >= c.modifyInterval {
		c.modifyNowPacketSize()
	}
	return
}

func NewClient(addr string) *Client {
	c :=  &Client{
		addr: addr,
		TimeoutLimit: 200 * time.Millisecond,
		nowPacketSize: PacketSize,
		retry: 3,
		modifyInterval: 100,
	}
	return c
}

func NewClientModify(addr string, modifyInterval int) *Client {
	c :=  &Client{
		addr: addr,
		TimeoutLimit: 200 * time.Millisecond,
		nowPacketSize: PacketSize,
		retry: 3,
		modifyInterval: modifyInterval,
	}
	return c
}