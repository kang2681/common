package drivers

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

func TestNewRedisConn(t *testing.T) {
	pool := NewRedisConn("127.0.0.1:6379", 10, 5, 10)
	conn := pool.Get()
	str, err := redis.String(conn.Do("SET", "test", "kang"))
	logrus.Infof("str:%+v，err：%+v", str, err)
	if err != nil {
		t.Error(err)
	}

	str, err = redis.String(conn.Do("GET", "test"))
	logrus.Infof("str:%+v，err：%+v", str, err)
	if err != nil {
		t.Error(err)
	}

	num, err := redis.Int64(conn.Do("DEL", "test"))
	logrus.Infof("num:%+v，err：%+v", num, err)
	if err != nil {
		t.Error(err)
	}
}
