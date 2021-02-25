package SendMsg

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
)

type SendClientPool struct {
	cpo []*Client
	workCh chan int
}

func (s *SendClientPool) SendTo(msg *chatMsg.Msg) {
	var id int
	select {
	case id = <- s.workCh:
		break
	}
	go func(idt int) {
		err := s.cpo[idt].SendTo(msg)
		if err != nil {
			log.Println(err)
		}
		s.workCh <- idt
	}(id)
}

func NewSendClientPool(sz int, addr string) *SendClientPool {
	ret := &SendClientPool {
		cpo: make([]*Client, sz),
		workCh: make(chan int, sz),
	}
	for i := 0 ; i < sz ; i ++ {
		ret.cpo[i] = NewClient(addr)
		ret.workCh <- i
	}
	return ret
}