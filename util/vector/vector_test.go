package vector

import (
	"log"
	"testing"
	"time"
)

func TestVector(t *testing.T) {
	o := make([]interface{}, 1)
	o = append(o, 1)
	o = append(o, 1)
	o = append(o, 1)
	o = append(o, 1)
	log.Println(cap(o))
}

func TestImpl_PushBack(t *testing.T) {
	var o = New()
	for i := 0; i < 20; i++ {
		o.PushBack(i)
	}
	for i := 0; i < 20; i++ {
		log.Println(o.Get(i))
	}
}

func BenchmarkImpl_Get(t *testing.B) {
	var o = New()
	start := time.Now()
	for i := 0; i < 100000; i++ {
		o.PushBack(i)
	}
	for i := 0; i < 29224; i++ {
		o.PopFront()
	}
	for i := 0; i < 2000; i++ {
		o.PopBack()
	}
	for i := 0; i < 2000000; i++ {
		o.PushFront(i)
	}
	for i := 0; i < 1000000; i++ {
		o.PopBack()
	}
	log.Println(time.Now().Sub(start))
}
