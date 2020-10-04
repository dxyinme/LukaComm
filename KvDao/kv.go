package KvDao

import "time"

const (
	KvTimeOut = 5 * time.Second
)

type KV interface {
	Get(string) ([]byte,error)
	Set(string, []byte) error
}
