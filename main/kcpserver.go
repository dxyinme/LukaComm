package main

import (
	"crypto/sha1"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
	"log"
)

func main() {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	if listener, err := kcp.ListenWithOptions("127.0.0.1:12345", block, 10, 1); err == nil {
		for {
			s, err := listener.AcceptKCP()
			if err != nil {
				log.Fatal(err)
			}
			go handleEcho(s)
		}
	} else {
		log.Fatal(err)
	}
}

// handleEcho send back everything it received
func handleEcho(conn *kcp.UDPSession) {
	defer func() {
		conn.Close()
		log.Println("end")
	}()
	for {
		buf := make([]byte, 32)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		if string(buf[:n]) == "end" {
			break
		}
		n, err = conn.Write([]byte("echo" + string(buf[:n])))
		if err != nil {
			log.Println(err)
			return
		}
	}

}