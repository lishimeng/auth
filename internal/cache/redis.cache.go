package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func Init(ctx context.Context, addr, password string, db int) {

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		client = nil
	}
	go Destroy(ctx)
}

func Get(key string) (value string, err error) {

	if client == nil {
		err = fmt.Errorf("cache is unuseable")
		return
	}

	value, err = client.Get(key).Result()
	return
}

func Set(key, value string, expire time.Duration) (err error) {

	if client == nil {
		err = fmt.Errorf("cache is unuseable")
		return
	}
	err = client.Set(key, value, expire).Err()
	return
}

func Destroy(ctx context.Context) {

	for {
		select {
		case <- ctx.Done():
			if client == nil {
				return
			}

			err := client.Close()
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}
