package KvDao

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

func TestNewRedisKv(t *testing.T) {
	testPool := NewRedisKv("127.0.0.1:6379")
	defer testPool.Close()
	err := testPool.Set("test", "just test")
	if err != nil {
		t.Fatal("test 1 : ", err)
	}
	res, err := testPool.Get("test1")
	if res != nil {
		t.Fatal("test 2 : res must be nil")
	}
	res, err = redis.String(testPool.Get("test"))
	if res != "just test" {
		t.Fatalf("test 3 : res not equal to %s", "just test")
	} else {
		log.Println(res)
	}
	err = testPool.Delete("test")
	if err != nil {
		t.Fatal(err)
	}
	err = testPool.Insert("xxx", "test")
	if err != nil {
		t.Fatal(err)
	}
	err = testPool.Insert("xxx", "test2")
	if err != nil {
		t.Fatal(err)
	}
	err = testPool.Insert("xxx", "test3")
	if err != nil {
		t.Fatal(err)
	}
	var setNow []string
	setNow, err = redis.Strings(testPool.GetMembers("xxx"))
	if err != nil {
		t.Fatal(err)
	}
	for _,v := range setNow {
		log.Println(v)
	}
	err = testPool.Remove("xxx", "test")
	if err != nil {
		t.Fatal(err)
	}
}
