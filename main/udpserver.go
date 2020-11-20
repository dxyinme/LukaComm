package main

import (
	"log"
	"net"
)

func main() {
	var (
		err error
	)
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		message := make([]byte, 8192)
		_, rAddr, err := conn.ReadFromUDP(message)
		go func(msg []byte) {
			if err != nil {
				log.Println(err)
				return
			}
			_, err := conn.WriteToUDP(append([]byte("echo"), msg...), rAddr)
			if err != nil {
				log.Println(err)
			}
		}(message)
	}

}
