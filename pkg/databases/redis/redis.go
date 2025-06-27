package redis

import (
	"os"

	"github.com/Nattakornn/cache/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type IRedisDb interface {
}

type redisDb struct {
	client *redis.Client
}

func ConnectRedisDb() IRedisDb { // TODO ADD Config for init redis
	opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	if err != nil {
		logger.Logger.Errorf("connect to redis failed: %v\n", err)
		os.Exit(1)
	}

	client := redis.NewClient(opt)
	return &redisDb{
		client: client,
	}
}
