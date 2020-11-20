package util

import (
	"bytes"
	"encoding/binary"
	"log"
	"testing"
)

func TestTransInt16(t *testing.T) {
	for i := 0; i < (1 << 8); i++ {
		int16i := int16(i)
		if int16i != ByteToInt16(Int16ToByte(int16i)) {
			t.Errorf("%d is trans error", int16i)
		} else {
			log.Printf("%d is verify OK", int16i)
		}
	}
}

func TestTranString(t *testing.T) {
	testStringList := []string{"sss", "31", "000000007000"}
	for i := range testStringList {
		nows := testStringList[i]
		if nows != ByteToString(StringToByteStaticLength(nows, 32)) {
			t.Errorf("No. %d test case is error, src : %s\n", i, nows)
		}
	}
}

func TestByteToUint64(t *testing.T) {
	for tCase := 0; tCase < 100; tCase++ {
		var now uint64
		now = uint64(tCase)
		buf := bytes.NewBuffer([]byte{})
		err := binary.Write(buf, binary.BigEndian, now)
		if err != nil {
			t.Errorf("error in %d, err is %v ", tCase, err)
		}
		//log.Println(buf.Bytes())
		res, err := ByteToUint64(buf.Bytes())
		if err != nil {
			t.Errorf("error in %d, err is %v ", tCase, err)
		}
		if res != now {
			t.Errorf("wanted %d but %d ", now, res)
		}
	}
}

func TestUint64ToByte(t *testing.T) {
	v := uint64(33423)
	ret := Uint64ToByte(v)
	rrt,_ := ByteToUint64(ret)
	log.Println(v)
	if rrt != v {
		t.Errorf("no equal")
	}
}