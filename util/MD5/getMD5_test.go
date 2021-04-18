package MD5

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
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

	t1 := make([]byte, 13)
	t2 := make([]byte, 25)
	nowt := []byte("wer  owr")
	copy(t1,nowt)
	copy(t2,nowt)
	if CalcMD5(t1) == CalcMD5(t2) {
		t.Errorf("%v == %v", t1, t2)
	}
}

func TestCalcMD5File(t *testing.T) {
	filename := "oo.txt"
	fp, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0 ; i < 1000; i ++ {
		_,_ = fp.Write([]byte("13kkewfiuwifgdew"))
	}
	fp.Close()
	md5res , err := CalcMD5File(filename)
	if err != nil {
		t.Fatal(err)
	}
	if md5res != "c3234aac44c95ea89747ef3d66946704" {
		t.Fatal("md5sum error")
	}
	err = os.Remove(filename)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCalcMD5ToByte(t *testing.T) {
	test1 := []byte("12hri3r ffioggrewqwwekppef")
	res := CalcMD5ToByte(test1)
	log.Println(res, len(res))
}