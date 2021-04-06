package TestPb

import (
	"encoding/json"
	"github.com/dxyinme/LukaComm/util"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
	"time"
)

var t1 = &BenchBody{
	T1:            "abbbbcddc",
	T2:            "assadswww",
	T3:            []byte("10103eig23r6r8 wfefwefjeoifj"),
	T4:            111223,
	T5:            3434583748,
	Arr1:          []int32{3,1,2,3,5,6},
	Arr2:          []string{"d1", "d2", "d3"},
	ReplyArr:      make([]*TestMsgReply, 10),
}

var resProtoG, resJsonG []byte

func thisInit() {
	resProtoG, _ = proto.Marshal(t1)
	resJsonG, _ = util.IJson.Marshal(t1)
}

func TestBenchBody_protoDecode(t *testing.T) {
	thisInit()
	var t1reborn = &BenchBody{}
	ti := time.Now()
	for i := 0 ; i < 10000; i ++ {
		_ = proto.Unmarshal(resProtoG, t1reborn)
	}
	log.Println(time.Now().Sub(ti).Milliseconds(), "ms")
}

func TestBenchBody_jsonDecode(t *testing.T) {
	thisInit()
	var t1reborn = &BenchBody{}
	ti := time.Now()
	for i := 0 ; i < 10000; i ++ {
		_ = json.Unmarshal(resJsonG, t1reborn)
	}
	log.Println(time.Now().Sub(ti).Milliseconds(), "ms")
}

func BenchmarkBenchBody_protoEncode(b *testing.B) {

	var (
		err error
		resProto []byte
	)

	resProto,err = proto.Marshal(t1)
	if err != nil {
		b.Fatal(err)
	}
	log.Printf("length proto : %d", len(resProto))
}

func BenchmarkBenchBody_jsonEncode(b *testing.B) {
	var (
		err error
		resJson []byte
	)

	resJson, err = util.IJson.Marshal(t1)

	if err != nil {
		b.Fatal(err)
	}

	log.Printf("length json : %d", len(resJson))
}