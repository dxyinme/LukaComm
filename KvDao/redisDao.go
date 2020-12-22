package KvDao

import (
	"errors"
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

func (rDao *RedisDao) Get(key string) (interface{}, error) {
	conn := rDao.pool.Get()
	defer conn.Close()
	var (
		res interface{}
		err error
	)
	res, err = conn.Do("GET", key)
	return res, err
}

func (rDao *RedisDao) Set(key string, value interface{}) error {
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

// Set operation

// Insert
func (rDao *RedisDao) Insert(setName string, item interface{}) error {
	conn := rDao.pool.Get()
	defer conn.Close()
	reply, err := conn.Do("SADD", setName, item)
	if reply == 0 {
		err = errors.New("item has existed")
	}
	return err
}

func (rDao *RedisDao) Remove(setName string, item interface{}) error {
	conn := rDao.pool.Get()
	defer conn.Close()
	reply, err := conn.Do("SREM", setName, item)
	if reply == 0 {
		err = errors.New("item has not existed")
	}
	return err
}

func (rDao *RedisDao) GetMembers(setName string) (ret []interface{}, err error) {
	conn := rDao.pool.Get()
	defer conn.Close()
	var reply interface{}
	reply, err = conn.Do("SMEMBERS", setName)
	if reply == nil {
		err = errors.New("setName has not existed")
	}
	ret = reply.([]interface{})
	return
}

func (rDao *RedisDao) Close() error {
	if rDao.pool == nil {
		return fmt.Errorf("have no pool")
	}
	return rDao.pool.Close()
}
