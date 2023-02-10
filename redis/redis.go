package redis

import (
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
)

var Pool *redis.Pool

func SetUp() {
	// creat conn pool
	Pool = &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", config.Redis.Host+":"+config.Redis.Port)
			if err != nil {
				return nil, err
			}
			conn.Do("Auth", config.Redis.Password)
			return conn, nil
		},
		// check wether connection available
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("Ping")
			return err
		},
	}
}

func Set(conn redis.Conn, key string, val interface{}) error {
	_, err := conn.Do("SET", key, val)
	if err != nil {
		return err
	}
	return nil
}

func Get(conn redis.Conn, key string) (string, error) {
	val, err := conn.Do("GET", key)
	str := string(val.([]byte))
	if err != nil {
		return "", err
	}
	return str, nil
}

func HSet(conn redis.Conn, key, field string, val interface{}) error {
	_, err := conn.Do("HSET", key, field, val)
	if err != nil {
		return err
	}
	return nil
}

func HGet(conn redis.Conn, key, field string) (string, error) {
	val, err := conn.Do("HGET", key, field)
	if err != nil {
		return "", err
	}
	return string(val.([]byte)), nil
}

func HGetAll(conn redis.Conn, key string) (map[string]string, error) {
	valsIf, err := conn.Do("HGETALL", key)
	if err != nil {
		log.Err(err).Msg("Faile to HGetAll")
		return nil, err
	}
	vals := valsIf.([]interface{})
	m := map[string]string{}

	for i := 0; i < len(vals); i += 2 {
		field := string(vals[i].([]byte))
		val := string(vals[i+1].([]byte))
		m[field] = val
	}
	return m, err
}

func Expire(conn redis.Conn, key string, expire time.Duration) error {
	_, err := conn.Do("EXPIRE", key, expire.Seconds())
	if err != nil {
		return err
	}
	return nil
}

func Exists(conn redis.Conn, key string) (bool, error) {
	result, err := conn.Do("EXISTS", key)
	if err != nil {
		return false, err
	}

	exists, err := strconv.ParseBool(strconv.Itoa(int(result.(int64))))
	if err != nil {
		return false, nil
	}
	return exists, nil
}
