package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/entity"
	"hh-go-bot/pkg/logger"
	"sync"
	"time"
)

func NewRedisClient(addr string, pwd string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
	})
}

type RedisService struct {
	redis redis.Client
	mu    *sync.Mutex
}

func NewRedisService(client *redis.Client) RedisService {
	return RedisService{
		redis: *client,
		mu:    &sync.Mutex{},
	}
}

func (s RedisService) Set(v any) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	vacancy, ok := v.(entity.Vacancy)
	if !ok {
		logger.Log.Warn("cannot use this value", "type", ok)
	}

	value, err := json.Marshal(vacancy)
	if err != nil {
		logger.Log.Warn("cannot marshal value", "vacancy", value)
	}

	status := s.redis.Set(ctx, vacancy.Id, value, time.Duration(config.All.Redis.Timeout)*time.Hour)
	if status.Err() != nil {
		logger.Log.Warn("key and value not added to cache", vacancy.Id, status.Err())
	} else {
		logger.Log.Debug("key and value added to cache", "id", vacancy.Id)
	}
}

func (s RedisService) Get(ctx context.Context, key string) error {
	strComm := s.redis.Get(ctx, key)
	return strComm.Err()
}
