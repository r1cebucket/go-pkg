package pkg_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
	"github.com/r1cebucket/gopkg/redis"
)

func init() {
	log.Setup("debug")
	config.Parse("../configs/conf.json")
	redis.SetUp()
}

func TestSet(t *testing.T) {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()

	var key string
	var err error

	key = "test:str"
	val := "test_val"
	err = redis.Set(redisConn, key, val)
	if err != nil {
		log.Err(err).Msg("Faile to set key")
		t.Error()
	}
}

func TestGet(t *testing.T) {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()

	var key string
	var err error

	key = "test:str"
	valTrue := "test_val"

	err = redis.Set(redisConn, key, valTrue)
	if err != nil {
		log.Err(err).Msg("Faile to set key")
		t.Error()
	}

	val, err := redis.Get(redisConn, key)
	if err != nil {
		log.Err(err).Msg("Faile to set key")
		t.Error()
	}
	if val != valTrue {
		t.Error()
	}
}

func TestHSet(t *testing.T) {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()

	var key, field string
	var err error

	key = "test:hash"
	field = "test_field"
	valTrue := "test_val"
	err = redis.HSet(redisConn, key, field, valTrue)
	if err != nil {
		log.Err(err).Msg("Faile to set hashmap")
		t.Error()
	}
	val, err := redis.HGet(redisConn, key, field)
	if err != nil {
		log.Err(err).Msg("Faile to get hashmap")
		t.Error()
	}
	if val != valTrue {
		t.Error()
	}
}

func TestHGet(t *testing.T) {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()

	var key, field, val string
	var err error

	key = "test:hash"
	field = "test_field"
	strTrue := "test_val"
	val, err = redis.HGet(redisConn, key, field)
	if err != nil {
		log.Err(err).Msg("Faile to get hashmap")
		t.Error()
	}
	if !(val == strTrue) {
		t.Error()
	}
}

func TestHGetAll(t *testing.T) {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()

	key := "test:keyname"
	vals, err := redis.HGetAll(redisConn, key)
	if err != nil {
		log.Err(err).Msg("Faile to get hashmap")
		t.Error()
	}
	log.Info().Msg(fmt.Sprint(vals))
}

func TestExpire(t *testing.T) {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()

	var key string
	var err error

	key = "test:expire"
	val := "test_val"
	err = redis.Set(redisConn, key, val)
	if err != nil {
		log.Err(err).Msg("Faile to set key")
		t.Error()
	}
	err = redis.Expire(redisConn, key, 1*time.Second)
	if err != nil {
		log.Err(err).Msg("Faile to set expire time")
	}
	time.Sleep(2 * time.Second)
	exists, err := redis.Exists(redisConn, key)
	if err != nil {
		log.Err(err).Msg("Faile to get exists")
	}
	if exists {
		t.Error()
	}
}

func TestExist(t *testing.T) {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()

	var key string
	var err error
	var exists bool

	key = "test:exists"
	err = redis.Set(redisConn, key, "exist")
	if err != nil {
		t.Error()
	}
	exists, err = redis.Exists(redisConn, key)
	if err != nil {
		log.Err(err).Msg("Faile to get result")
		t.Error()
	}
	if !exists {
		t.Error()
	}

	key = "test:not_exists"
	exists, err = redis.Exists(redisConn, key)
	if err != nil {
		log.Err(err).Msg("Faile to get result")
		t.Error()
	}
	if exists {
		t.Error()
	}
}
