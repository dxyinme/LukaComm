package KvDao

import (
	"log"
	"testing"
)

func TestNewRedisKv(t *testing.T) {
	testPool := NewRedisKv("yourHost")
	defer testPool.Close()
	err := testPool.Set("test", []byte("just test"))
	if err != nil {
		t.Fatal("test 1 : ", err)
	}
	res, err := testPool.Get("test1")
	if res != nil {
		t.Fatal("test 2 : res must be nil")
	}
	res, err = testPool.Get("test")
	if string(res) != "just test" {
		t.Fatalf("test 3 : res not equal to %s", "just test")
	} else {
		log.Println(string(res))
	}
	err = testPool.Delete("test")
	if err != nil {
		t.Fatal(err)
	}
}
