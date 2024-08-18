package gredis

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/AOaixuexi/go-gin-example-follow/pkg/setting"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
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

// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
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


// Setup initializes the Redis connection pool based on the configuration settings.
// It returns an error if there is any issue with the initialization.
// func Setup() error {}

// Set sets the value of a key in Redis with the specified data and expiration time.
// It takes the key as a string, the data as an interface{}, and the time as an integer.
// It returns an error if there is any issue with setting the value.
// func Set(key string, data interface{}, time int) error {}

// Exists checks if a key exists in Redis.
// It takes the key as a string and returns a boolean indicating whether the key exists or not.
// func Exists(key string) bool {}

// Get retrieves the value of a key from Redis.
// It takes the key as a string and returns the value as a byte slice.
// It returns an error if there is any issue with retrieving the value.
// func Get(key string) ([]byte, error) {}

// Delete deletes a key from Redis.
// It takes the key as a string and returns a boolean indicating whether the key was deleted or not.
// It returns an error if there is any issue with deleting the key.
// func Delete(key string) (bool, error) {}

// LikeDeletes deletes all keys in Redis that match a given pattern.
// It takes the pattern as a string and returns an error if there is any issue with deleting the keys.
// func LikeDeletes(key string) error {}