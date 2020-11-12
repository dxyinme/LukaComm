package KvDao

import "time"

const (
	KvTimeOut = 5 * time.Second
)

type KV interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Exists(key string) (bool, error)
	Delete(key string) error
}
