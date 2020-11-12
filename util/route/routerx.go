package route

import (
	"math/rand"
	"sync"
)

type RouterX interface {
	GetRoute() EndPoint
	ReBuild(epList []EndPoint)
}

type Impl struct {
	endPointMap       map[int]EndPoint
	weightPreAddArray []int
	mutex             sync.Mutex
}

func (im *Impl) GetRoute() EndPoint {
	im.mutex.Lock()
	defer im.mutex.Unlock()
	v := rand.Intn(im.weightPreAddArray[len(im.weightPreAddArray)-1])
	l, r, ans := 1, len(im.weightPreAddArray), 1
	for l <= r {
		m := (l + r) >> 1
		if v <= im.weightPreAddArray[m-1] {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return im.endPointMap[ans-1]
}

func (im *Impl) ReBuild(epList []EndPoint) {
	im.mutex.Lock()
	defer im.mutex.Unlock()
	im.endPointMap = make(map[int]EndPoint)
	im.weightPreAddArray = make([]int, len(epList))
	for i := 0; i < len(epList); i++ {
		im.endPointMap[i] = epList[i]
		if i == 0 {
			im.weightPreAddArray[i] = epList[i].Weight
		} else {
			im.weightPreAddArray[i] = epList[i].Weight + im.weightPreAddArray[i-1]
		}
	}
}

func NewRouteX(epList []EndPoint) RouterX {
	Tmp := &Impl{}
	Tmp.ReBuild(epList)
	return Tmp
}
