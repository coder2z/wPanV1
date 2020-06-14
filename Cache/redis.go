package Redis

import (
	"encoding/json"
	"time"
	"wPan/v1/Config"

	"github.com/gomodule/redigo/redis"
)

var Conn *redis.Pool

func InitRedis() error {
	Conn = &redis.Pool{
		MaxIdle:     Config.RedisSetting.MaxIdle,
		MaxActive:   Config.RedisSetting.MaxActive,
		IdleTimeout: Config.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Config.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if Config.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", Config.RedisSetting.Password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) (bool, error) {
	conn := Conn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	reply, err := redis.String(conn.Do("SET", key, value))
	_, _ = conn.Do("EXPIRE", key, time)
	if reply == "OK" {
		return true, err
	}
	return false, err
}

func Exists(key string) bool {
	conn := Conn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := Conn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := Conn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := Conn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
