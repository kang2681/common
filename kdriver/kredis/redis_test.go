package kredis

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestNewRedisConn(t *testing.T) {
	pool := NewRedisConn("127.0.0.1:6379", 10, 5, 10)
	conn := pool.Get()
	_, err := redis.String(conn.Do("SET", "test", "kang"))
	if err != nil {
		t.Error(err)
	}

	_, err = redis.String(conn.Do("GET", "test"))
	if err != nil {
		t.Error(err)
	}

	_, err = redis.Int64(conn.Do("DEL", "test"))
	if err != nil {
		t.Error(err)
	}
}
