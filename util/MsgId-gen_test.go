package util

import (
	"log"
	"testing"
)

func TestMsgIdGen(t *testing.T) {
	var (
		UID          string = "haha"
		registerTime int = 5
		lastMsgTime  int = 10
		tip          int = 0
		res 		 string = ""
	)
	lastMsgTime,tip,res = MsgIdGen(UID, registerTime, lastMsgTime, tip)
	log.Println(res)
}