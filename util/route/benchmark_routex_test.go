package route

import (
	"testing"
)

func BenchmarkImpl_GetRoute(b *testing.B) {
	r:= NewRouteX([]EndPoint{
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
	b.Run("100000" , func(b *testing.B) {
		etm := "Simon6"
		for i := 0 ; i < 100000 ; i ++ {
			v := r.GetRoute().Name
			if etm == v {

			}
		}
	})

}