package board

import (
	"encoding/json"
	"time"
	"github.com/garyburd/redigo/redis"
)

func newPool(config *ServerConfig) *redis.Pool{
	return &redis.Pool {
		MaxIdle : 4,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis.Address)
			if err != nil {
				return nil, err
			}
			if config.Redis.Auth != "" {
				if _, err := c.Do("AUTH", config.Redis.Auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", config.Redis.DB); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func getPostId(r *redis.Redis) (string, error) {
	res, err := redis.String(r.Do("INCR", "PostIdSeq"))
	return res, err
}

func getUserId(r, *redis.Redis) (string, error) {
	res, err := redis.String(r.Do("INCR", "UserIdSeq"))
	return res, err
}


