package main

import (
	"log"
	"net"
	"sync"
	"time"
)

func main()  {
	startTime := time.Now()
	var reqCount int = 10000
	var wg sync.WaitGroup
	wg.Add(reqCount)
	for i := 0 ; i < reqCount ; i ++ {
		go func(name string) {
			conn, err := net.Dial("udp", "127.0.0.1:8080")
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
			_, err = conn.Read(resp)
			if err != nil {
				log.Println(err)
				return
			}
			Len := 0
			for ;resp[Len] != 0; {
				Len ++
			}
			respString := string(resp[:Len])
			if respString != "echo" + name {
				log.Printf("err Echo , resp is [%s] , wanted [%s]", respString, "echo" + name)
				log.Printf("resp %v", []byte(respString))
				log.Printf("want %v", []byte("echo" + name))
			}
			log.Println("OK")
			wg.Done()
		}("UUUUU")
	}
	wg.Wait()
	log.Printf("use time : %v ms", time.Now().Sub(startTime).Milliseconds())
}
