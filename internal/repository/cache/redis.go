package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hh-go-bot/internal/config"
	"time"
)

func NewRedisClient(addr string, pwd string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr, // адрес вашего Redis-сервера
		Password: pwd,  // пароль, если он установлен
		DB:       0,    // индекс базы данных
	})
}

type RedisService struct {
	client *redis.Client
}

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{client: client}
}

func (s RedisService) ConvertAndSet(ctx context.Context, key string, value any) {
	timeout := config.All.Redis.Timeout * time.Hour
	s.client.Set(ctx, key, value, timeout)
}
