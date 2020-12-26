package main

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var reqCount int = 2500
	var wg sync.WaitGroup
	wg.Add(reqCount)
	var timeoutCnt = 0

	for i := 0; i < reqCount; i++ {
		go func(name string) {
			conn, err := net.Dial("udp", "localhost:8080")
			if err != nil {
				log.Println(err)
				return
			}
			defer conn.Close()
			_, err = conn.Write([]byte(name))
			if err != nil {
				log.Println(err)
				return
			}
			resp := make([]byte, 8192)
			err = func(b []byte) error {
				ch := make(chan bool, 1)
				go func() {
					_, err = conn.Read(b)
					close(ch)
				}()
				select {
				case <-ch:
					return nil
				case <-time.After(20 * time.Second):
					return errors.New("timeout")
				}
			}(resp)
			if err != nil {
				log.Println(err)
				wg.Done()
				timeoutCnt++
				return
			}
			Len := 0
			for resp[Len] != 0 {
				Len++
			}
			respString := string(resp[:Len])
			if respString != "echo"+name {
				log.Printf("err Echo , resp is [%s] , wanted [%s]", respString, "echo"+name)
				log.Printf("resp %v", []byte(respString))
				log.Printf("want %v", []byte("echo"+name))
			}
			log.Println("OK")
			wg.Done()
		}("UUUUU")
	}
	wg.Wait()
	log.Printf("use time : %v ms in %v requests, timeoutCnt %v",
		time.Now().Sub(startTime).Milliseconds(), reqCount, timeoutCnt)
}
