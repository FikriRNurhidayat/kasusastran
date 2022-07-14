package redis

import "github.com/go-redis/redis/v8"

func ConnectRedis(dataSourceName string) (client *redis.Client, err error) {
	opt, err := redis.ParseURL(dataSourceName)

	if err != nil {
		return client, err
	}

	return redis.NewClient(opt), nil
}
