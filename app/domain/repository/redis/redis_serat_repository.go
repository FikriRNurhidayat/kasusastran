package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc/grpclog"
)

type RedisSeratRepository struct {
	redisRepository RedisRepository
	seratRepository repository.SeratRepository
	logger          grpclog.LoggerV2
}

const (
	REDIS_SERAT_KEY             = "serats"
	REDIS_SERAT_EXPIRATION_TIME = 5 * time.Minute
)

func (r *RedisSeratRepository) createKey(id uuid.UUID) string {
	return fmt.Sprintf(REDIS_SERAT_KEY+"/%s", id.String())
}

func (r *RedisSeratRepository) Create(ctx context.Context, i *entity.Serat) (o *entity.Serat, err error) {
	// Hit database
	o, err = r.seratRepository.Create(ctx, i)
	if err != nil {
		return nil, err
	}

	// Save to cache
	go r.redisRepository.Set(ctx, r.createKey(o.ID), o, REDIS_SERAT_EXPIRATION_TIME)

	return o, nil
}

// Delete implements repository.SeratRepository
func (r *RedisSeratRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	go r.redisRepository.Del(ctx, r.createKey(id))
	return r.seratRepository.Delete(ctx, id)
}

// Get implements repository.SeratRepository
func (r *RedisSeratRepository) Get(ctx context.Context, id uuid.UUID) (serat *entity.Serat, err error) {
	cacheKey := r.createKey(id)
	serat = &entity.Serat{}

	// Hit cache
	err = r.redisRepository.Get(ctx, cacheKey, serat)
	if err != nil {
		return nil, err
	}

	// Return it if there's value on cache
	if serat.ID != uuid.Nil {
		return serat, nil
	}

	// Hit database
	serat, err = r.seratRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Set Cache
	go r.redisRepository.Set(ctx, cacheKey, *serat, REDIS_SERAT_EXPIRATION_TIME)

	return serat, nil
}

// List implements repository.SeratRepository
func (r *RedisSeratRepository) List(ctx context.Context, query *repository.ListQuery) (serats []*entity.Serat, count uint32, err error) {
	return r.seratRepository.List(ctx, query)
}

// Update implements repository.SeratRepository
func (r *RedisSeratRepository) Update(ctx context.Context, id uuid.UUID, userat *entity.Serat) (serat *entity.Serat, err error) {
	// Hit database
	serat, err = r.seratRepository.Update(ctx, id, userat)
	if err != nil {
		return nil, err
	}

	// Set cache
	go r.redisRepository.Set(ctx, r.createKey(id), serat, REDIS_SERAT_EXPIRATION_TIME)

	return serat, nil
}

func NewRedisSeratRepository(redisRepository RedisRepository, seratRepository repository.SeratRepository, logger grpclog.LoggerV2) repository.SeratRepository {
	return &RedisSeratRepository{
		redisRepository: redisRepository,
		seratRepository: seratRepository,
		logger:          logger,
	}
}
