package KvDao

import "time"

const (
	KvTimeOut = 5 * time.Second
)

type KV interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Exists(key string) (bool, error)
	Delete(key string) error
}
