package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/entity"
	"hh-go-bot/pkg/logger"
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
}

func NewRedisService(client *redis.Client) RedisService {
	return RedisService{
		redis: *client,
	}
}

func (s RedisService) ConvertAndSet(ctx context.Context, v any) {
	vacancies, ok := v.(entity.Vacancies)
	if !ok {
		logger.Log.Warn("cannot use this value", "type", ok)
	}
	for _, vacancy := range vacancies.Items {
		value, err := json.Marshal(vacancy)
		if err != nil {
			logger.Log.Warn("cannot marshal value", "vacancy", value)
		}
		s.redis.Set(ctx, vacancy.Id, value, time.Duration(config.All.Redis.Timeout)*time.Hour)
		logger.Log.Debug("key and value added to cache", "id", vacancy.Id)
	}
}

func (s RedisService) Get(ctx context.Context, key string) error {
	strComm := s.redis.Get(ctx, key)
	return strComm.Err()
}
