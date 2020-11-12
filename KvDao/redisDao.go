package KvDao

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisDao struct {
	pool      *redis.Pool
	redisHost string
}

func NewRedisKv(redisHost string) *RedisDao {
	return &RedisDao{
		pool: &redis.Pool{
			Dial: func() (conn redis.Conn, e error) {
				c, err := redis.Dial("tcp", redisHost)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:     3,
			IdleTimeout: KvTimeOut,
		},
		redisHost: redisHost,
	}
}

func (rDao *RedisDao) Get(key string) ([]byte, error) {
	conn := rDao.pool.Get()
	defer conn.Close()
	var (
		res []byte
		err error
	)
	res, err = redis.Bytes(conn.Do("GET", key))
	return res, err
}

func (rDao *RedisDao) Set(key string, value []byte) error {
	conn := rDao.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	return err
}

func (rDao *RedisDao) Exists(key string) (bool, error) {
	conn := rDao.pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}

func (rDao *RedisDao) Delete(key string) error {
	conn := rDao.pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}

func (rDao *RedisDao) Close() error {
	if rDao.pool == nil {
		return fmt.Errorf("have no pool")
	}
	return rDao.pool.Close()
}
