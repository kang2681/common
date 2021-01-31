package drivers

import "time"

//NewRedisConn 获取连接池 18.03.21更新
func NewRedisConn(server string, max, min, timeoutSecond int) *redis.Pool {
	timeout := time.Duration(timeoutSecond) * time.Second
	return &redis.Pool{
		MaxIdle:   min,
		MaxActive: max,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialTimeout("tcp", server, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
