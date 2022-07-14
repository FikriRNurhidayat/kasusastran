package config

import "fmt"

func GetRedisConnectionString() string {
	return fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		Getenv("REDIS_USER", ""),
		Getenv("REDIS_PASSWORD", ""),
		Getenv("REDIS_HOST", "localhost"),
		Getenv("REDIS_PORT", "6379"),
		Getenv("REDIS_NAME", "0"),
	)
}
