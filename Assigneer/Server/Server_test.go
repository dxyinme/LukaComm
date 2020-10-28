package Server

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestNormal(t *testing.T) {
	gs := []int{1,2,3,4,5}
	var wg sync.WaitGroup
	wg.Add(len(gs))
	for _,v := range gs {
		go func(t int) {
			time.Sleep(time.Millisecond * 50)
			log.Println(t)
			wg.Done()
		}(v)
	}
	wg.Wait()
}