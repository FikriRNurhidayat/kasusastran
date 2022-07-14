package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/grpclog"
)

type RedisRepository interface {
	Get(ctx context.Context, k string, o any) error
	Set(ctx context.Context, k string, v interface{}, e time.Duration) (err error)
	Del(ctx context.Context, k string) error
}

type redisRepository struct {
	client redis.UniversalClient
	logger grpclog.LoggerV2
}

// Del implements RedisRepository
func (r *redisRepository) Del(ctx context.Context, k string) error {
	cmd := r.client.Del(ctx, k)
	if err := cmd.Err(); err != nil {
		return err
	}

	r.logger.Infof("redis: delete cache: %s", k)
	return nil
}

func (r *redisRepository) Get(ctx context.Context, k string, o any) error {
	value, err := r.client.Get(ctx, k).Result()

	if err == redis.Nil {
		r.logger.Infof("redis: cache not found: %s", k)
		*&o = nil
		return nil
	}

	if err != nil {
		*&o = nil
		return err
	}

	r.logger.Infof("redis: parsing cache: %s", k)
	return json.Unmarshal([]byte(value), o)
}

func (r *redisRepository) Set(ctx context.Context, k string, v interface{}, e time.Duration) (err error) {
	value, err := json.Marshal(v)

	if err != nil {
		r.logger.Warning("redis: fail to write cache: %s", k)
		return err
	}

	r.logger.Infof("redis: writing cache: %s", k)
	return r.client.Set(ctx, k, string(value), e).Err()
}

func NewRedisRepository(client redis.UniversalClient, logger grpclog.LoggerV2) RedisRepository {
	return &redisRepository{
		client: client,
		logger: logger,
	}
}
