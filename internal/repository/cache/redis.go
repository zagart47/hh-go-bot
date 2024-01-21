package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/entity"
	"log"
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
	redis redis.Client
}

func NewRedisService(client *redis.Client) RedisService {
	return RedisService{
		redis: *client,
	}
}

func (s RedisService) ConvertAndSet(ctx context.Context, value any) {
	vacancies, ok := value.(entity.Vacancies)
	if !ok {
		log.Println("type error")
	}
	for _, v := range vacancies.Items {
		s.redis.Set(ctx, v.Id, v, time.Duration(config.All.Redis.Timeout)*time.Hour)
	}
}
