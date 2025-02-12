package main

import (
	"context"
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"github.com/redis/go-redis/v9"
)

func main() {
	rcg := config.GetRuntimeConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rcg.RedisConfig.RedisHost, rcg.RedisConfig.RedisPort),
		Password: rcg.RedisConfig.RedisPassword,
	})
	res := rdb.Expire(context.Background(), "test_zset", 10)
	fmt.Println(res.Val())
	fmt.Println(res.Result())
	fmt.Println(res.Err())
}
