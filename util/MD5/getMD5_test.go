package MD5

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func TestCalcMD5(t *testing.T) {
	test1 := &chatMsg.Msg{
		From:           "luka1",
		Target:         "luka2",
		Content:        []byte("UUUUUXXOS"),
		MsgType:        chatMsg.MsgType_Single,
		MsgContentType: chatMsg.MsgContentType_Text,
	}
	testPB, err := proto.Marshal(test1)
	if err != nil {
		log.Println(err)
	}
	md5 := CalcMD5(testPB)
	log.Println(md5)
}
