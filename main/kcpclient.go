package main

import (
	"crypto/sha1"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"log"
	"sync"
	"time"
)

func main() {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	startTime := time.Now()
	reqCount := 10000
	wg := sync.WaitGroup{}
	wg.Add(reqCount)
	sess, err := kcp.DialWithOptions("127.0.0.1:12345", block, 10, 1)
	if err != nil {
		log.Fatal(err)
	}
	send := func(v int, data string) {
		buf := make([]byte, len(data) + len("echo"))
		if _, err := sess.Write([]byte(data)); err == nil {

			if _, err := io.ReadFull(sess,buf); err == nil {
				log.Printf("%d ok : %s", v, string(buf))
			} else {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
	for i := 0 ; i < reqCount; i ++ {
		go func (v int) {
			send(v, "UUUUU")
			wg.Done()
		}(i)
	}
	wg.Wait()

	log.Printf("use time : %v ms in %v requests",
		time.Now().Sub(startTime).Milliseconds(), reqCount)
}
