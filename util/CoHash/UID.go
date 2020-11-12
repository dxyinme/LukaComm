package CoHash

import (
	"github.com/spaolacci/murmur3"
	"strconv"
	"sync"
)

const beginUID = 10000

var (
	NowUID uint64 = beginUID
	mu     sync.Mutex
)

type UID struct {
	Uid string
}

func (u *UID) GetHash() uint32 {
	return murmur3.Sum32([]byte(u.Uid))
}

func NewUID() *UID {
	mu.Lock()
	defer mu.Unlock()
	NowUID++
	return &UID{
		strconv.Itoa(int(NowUID)),
	}
}
