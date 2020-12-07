package main

import (
	"flag"
	"fmt"
	"github.com/dxyinme/LukaComm/CynicU/P2P"
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"strings"
)

var (
	listenerAddr = flag.String("listenerAddr", ":8080", "listener address")
	myUID = flag.String("myUID", "testUID", "UID")
)
func readLine() string {
	var (
		b []byte
		c byte
	)
	for {
		_,_ = fmt.Scanf("%c", &c)
		if c == '\n' || c == '\r' {
			return string(b)
		} else {
			b = append(b, c)
		}
	}
}

func main() {
	flag.Parse()
	p2p := P2P.NewP2P()
	err := p2p.Listen(*listenerAddr)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan *chatMsg.Msg, 10)
	go func() {
		for {
			select {
			case it := <-p2p.RecvMsg:
				ch <- it
			}
		}
	}()
	for {
		var (
			s string
			UID string
			)
		fmt.Print("$")
		s = readLine()
		s = strings.TrimSpace(s)
		if len(s) < 1 {
			continue
		} else if s[0] == 'S' {
			var (
				addr string
				content string
			)
			_,_ = fmt.Sscanf(s, "S=%s T=%s", &content, &addr)
			//log.Println(content, "~", addr)
			_ = p2p.Send(&chatMsg.Msg{
				From:           "",
				Target:         addr,
				Content:        []byte(content),
				MsgType:        chatMsg.MsgType_Single,
				MsgContentType: chatMsg.MsgContentType_Text,
				SendTime:       "",
				GroupName:      "",
				Spread:         false,
				MsgId:          "",
			}, addr)

		} else if s[0] == 'P' {
			chLen := len(ch)
			for i := 0 ; i < chLen ; i ++ {
				log.Println(string((<-ch).Content))
			}
		} else if s == "exit" {
			return
		} else if s[0] == 'G' {
			_,_ = fmt.Sscanf(s, "G=%s", &UID)
			addr := P2P.GetAddr(UID)
			log.Println(addr)
		} else if s[0] == 's' {
			err = P2P.SyncAddr(*myUID, strings.Split(*listenerAddr, ":")[1])
			if err != nil {
				log.Println(err)
			}
		}
	}
}
