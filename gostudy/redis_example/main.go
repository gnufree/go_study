package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func NewPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/*
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			*/
			return c, err

		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
func testSet() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("conn failed ,err:%v\n", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("set", "abc", 100)
	if err != nil {
		fmt.Printf("conn Do set failed,err:%v\n", err)
		return
	}
	r, err := redis.Int(conn.Do("get", "abc"))
	if err != nil {
		fmt.Printf("get key failed, err:%v\n", err)
	}
	fmt.Printf("get r:%v\n", r)
}

func testExpire() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("conn failed ,err:%v\n", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("expire", "abc", 10)
	if err != nil {
		fmt.Printf("conn Do set expire,err:%v\n", err)
		return
	}
}

func testHget() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("conn failed ,err:%v\n", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("Hset", "books", "abc", 100)
	if err != nil {
		fmt.Printf("conn Do set expire,err:%v\n", err)
		return
	}
	r, err := redis.Int(conn.Do("Hget", "books", "abc"))
	if err != nil {
		fmt.Printf("Hget books failed, err:%v\n", err)
		return
	}
	fmt.Printf("hset :%v\n", r)
}

func main() {
	// testSet()
	// testExpire()
	// testHget()
	pool = NewPool("127.0.0.1:6379", "")
	for {
		time.Sleep(time.Second)
		conn := pool.Get()
		conn.Do("set", "abc", 100)

		r, err := redis.Int(conn.Do("get", "abc"))
		if err != nil {
			fmt.Printf("do failed, err:%v\n", err)
			continue
		}

		fmt.Printf("get from redis, result:%v\n", r)
	}
}
