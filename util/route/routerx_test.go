package route

import (
	"log"
	"testing"
)

func TestImpl_GetRoute(t *testing.T) {
	r := NewRouteX([]EndPoint{
		{
			Ip:     "1.0.2.0",
			Port:   138,
			Weight: 15,
			Name:   "Usx",
		},
	})
	for i := 0; i < 10; i++ {
		log.Println(r.GetRoute().Name)
	}
	log.Println("-------------------")
	r.ReBuild([]EndPoint{
		{
			Ip:     "1.0.0.0",
			Port:   22392,
			Weight: 134,
			Name:   "Simon",
		},
		{
			Ip:     "1.0.2.0",
			Port:   22393,
			Weight: 15,
			Name:   "Simon1",
		},
		{
			Ip:     "1.0.2.0",
			Port:   22172,
			Weight: 15,
			Name:   "Simon2",
		},
		{
			Ip:     "1.0.2.0",
			Port:   22,
			Weight: 15,
			Name:   "Simon3",
		},
		{
			Ip:     "1.0.2.0",
			Port:   22399,
			Weight: 15,
			Name:   "Simon4",
		},
		{
			Ip:     "1.0.2.0",
			Port:   22394,
			Weight: 158,
			Name:   "Simon5",
		},
		{
			Ip:     "1.0.2.0",
			Port:   22395,
			Weight: 15,
			Name:   "Simon6",
		},
	})
	for i := 0; i < 10; i++ {
		log.Println(r.GetRoute().Name)
	}
}
