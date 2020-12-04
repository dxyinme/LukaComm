package util

import (
	"log"
	"testing"
)

func TestMsgIdGen(t *testing.T) {
	var (
		UID          string = "haha"
		registerTime int64 = 5
		lastMsgTime  int64 = 10
		tip          int64 = 0
		res 		 string = ""
	)
	lastMsgTime,tip,res = MsgIdGen(UID, registerTime, lastMsgTime, tip)
	log.Println(res)
}