package redis

import (
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rcg := config.GetRuntimeConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rcg.RedisConfig.RedisHost, rcg.RedisConfig.RedisPort),
		Password: rcg.RedisConfig.RedisPassword,
	})

}

func GetRedisClient() *redis.Client {
	return rdb
}
