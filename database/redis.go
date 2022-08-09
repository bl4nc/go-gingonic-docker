package database

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

const (
	redisExpire = 60 * 60
)

func RedisConnect() redis.Conn {
	fmt.Println(redisExpire)

	c, err := redis.Dial("tcp", "redis:6379",
		redis.DialPassword(os.Getenv("REDIS_PASS")),
	)
	if err != nil {
		panic(err)
	}
	return c
}

func SetRedis(key string, value []byte) error {
	conn := RedisConnect()
	defer conn.Close()
	_, err := conn.Do("SET", key, []byte(value))
	if err != nil {
		panic(err)
	}
	conn.Do("EXPIRE", key, redisExpire) //10 Minutes
	return err
}

func GetRedis(key string) ([]byte, error) {
	conn := RedisConnect()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}
