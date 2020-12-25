package TestPb

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

func BenchmarkBenchBody_protoEncode(b *testing.B) {
	t1 := &BenchBody{
		T1:            "abbbbcddc",
		T2:            "assadswww",
		T3:            []byte("10103eig23r6r8 wfefwefjeoifj"),
		T4:            111223,
		T5:            3434583748,
		Arr1:          []int32{3,1,2,3,5,6},
		Arr2:          []string{"d1", "d2", "d3"},
		ReplyArr:      make([]*TestMsgReply, 10),
	}
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
	t1 := &BenchBody{
		T1:            "abbbbcddc",
		T2:            "assadswww",
		T3:            []byte("10103eig23r6r8 wfefwefjeoifj"),
		T4:            111223,
		T5:            3434583748,
		Arr1:          []int32{3,1,2,3,5,6},
		Arr2:          []string{"d1", "d2", "d3"},
		ReplyArr:      make([]*TestMsgReply, 10),
	}
	var (
		err error
		resJson []byte
	)

	resJson, err = json.Marshal(t1)

	if err != nil {
		b.Fatal(err)
	}

	log.Printf("length json : %d", len(resJson))
}